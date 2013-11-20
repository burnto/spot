package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/exec"
	"strings"
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
	err = watcher.Watch(".")
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	// Process events
	for {
		cmd, err := startProcess(args)
		if err != nil {
			log.Println(err)
		}

		select {
		case ev := <-watcher.Event:
			log.Println("spotted", ev)
			if cmd != nil {
				cmd.Process.Kill()
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
		log.Println("Building", exe)
		cmd = exec.Command("go", "build", exe)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		exe = exe[:len(exe)-len(".go")]
	}

	cmd = exec.Command(exe, args[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
