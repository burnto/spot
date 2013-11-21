package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/fsnotify"
	"github.com/kr/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: spot <command>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(2)
	}

	// Watch the current directory
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}

	walker := fs.Walk(".")
	for walker.Step() {
		if err := walker.Err(); err != nil {
			log.Fatalln(err)
		}
		dir := filepath.Base(walker.Path())
		if walker.Stat().IsDir() {
			if dir != "." && (strings.HasPrefix(dir, "_") || strings.HasPrefix(dir, ".")) {
				walker.SkipDir()
				continue
			}
			err = watcher.Watch(walker.Path())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	defer watcher.Close()

	// Process events
	for {
		cmd, err := startProcess(args)
		if err != nil {
			log.Fatalln(err)
		}

		select {
		case ev := <-watcher.Event:
			events := []*fsnotify.FileEvent{ev}
			// spin for a moment to try to batch up other events
			timeout := time.NewTimer(time.Millisecond * 50)
		wait:
			for {
				select {
				case ev := <-watcher.Event:
					events = append(events, ev)
				case err := <-watcher.Error:
					log.Println("error", err)
				case <-timeout.C:
					break wait
				}
			}
			log.Println("spotted", events)
			if cmd != nil {
				err := cmd.Process.Kill()
				if err != nil {
					log.Fatalln(err)
				}
				_, err = cmd.Process.Wait()
				if err != nil {
					log.Fatalln(err)
				}
			}
		case err := <-watcher.Error:
			log.Println("error", err)
		}
	}
}

func startProcess(args []string) (*exec.Cmd, error) {
	var cmd *exec.Cmd

	// Build if it's a go file
	exe := args[0]
	if strings.HasSuffix(exe, ".go") {
		target := path.Join(string(os.PathSeparator)+"tmp", "spotted-"+exe[:len(exe)-len(".go")])

		log.Println("Building", exe)
		cmd = exec.Command("go", "build", "-o", target, exe)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		exe = target
	}

	cmd = exec.Command(exe, args[1:]...)
	log.Println("Running", strings.Join(args, " "))
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
