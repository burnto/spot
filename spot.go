package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/exec"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: spot <command>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func startProcess(args []string) *exec.Cmd {

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	return cmd
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(2)
	}
	cmd := startProcess(args)

	// Set up fsnotify
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Watch the current directory
	err = watcher.Watch(".")
	if err != nil {
		log.Fatal(err)
	}

	// Process events
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
			cmd.Process.Kill()
			cmd = startProcess(args)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
