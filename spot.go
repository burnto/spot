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

var watcher *fsnotify.Watcher

func usage() {
	fmt.Fprintf(os.Stderr, "usage: spot <command>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	defer watcher.Close()

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(2)
	}

	cmd := startProcess(args)

	// Process events
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("spotted", ev)
			cmd.Process.Kill()
			cmd = startProcess(args)
		case err := <-watcher.Error:
			log.Println("error", err)
		}
	}
}

func startProcess(args []string) *exec.Cmd {

	if watcher != nil {
		watcher.Close()
		watcher = nil
	}

	var cmd *exec.Cmd
	var err error

	// Build if it's a go file
	exe := args[0]
	if strings.HasSuffix(exe, ".go") {
		log.Println("Building", exe)
		cmd = exec.Command("go", "build", exe)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		exe = exe[:-3]
	}

	// Set up fsnotify
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Watch the current directory
	err = watcher.Watch(".")
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}
