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

	// Process events
	for {
		cmd, _ := startProcess(args)
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

	if watcher != nil {
		watcher.Close()
		watcher = nil
	}

	var cmd *exec.Cmd
	var cmdErr, watcherErr error

	// Build if it's a go file
	exe := args[0]
	if strings.HasSuffix(exe, ".go") {
		log.Println("Building", exe)
		cmd = exec.Command("go", "build", exe)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmdErr = cmd.Run()
		exe = exe[:len(exe)-3]
	}

	// Set up fsnotify
	watcher, watcherErr = fsnotify.NewWatcher()
	if watcherErr != nil {
		log.Panic(watcherErr)
	}

	// Watch the current directory
	watcherErr = watcher.Watch(".")
	if watcherErr != nil {
		log.Panic(watcherErr)
	}
	if cmdErr != nil {
		return nil, cmdErr
	}

	cmd = exec.Command(exe, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmdErr = cmd.Start()
	if cmdErr != nil {
		return nil, cmdErr
	}
	return cmd, nil
}
