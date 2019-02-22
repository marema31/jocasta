package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/marema31/jocasta/config"
	"github.com/marema31/jocasta/logwriter"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Println("usage: jocasta command to run with args")
		os.Exit(0)
	}

	config, err := config.New(".", "jocasta", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	logerr, err := logwriter.New("err", config)
	if err != nil {
		log.Fatal(err)
	}

	logout, err := logwriter.New("out", config)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdout and stderr
	// wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(2)
	go transfer(stdout, logout, &wg)
	go transfer(stderr, logerr, &wg)

	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Unable to exec %s: %s\n", os.Args[1], err)
	}
}

func transfer(in io.ReadCloser, out io.Writer, wg *sync.WaitGroup) {
	if _, err := io.Copy(out, in); err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
