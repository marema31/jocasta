package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/marema31/jocasta/config"
	"github.com/marema31/jocasta/logwriter"
)

func main() {

	i := 1
	configPath := "."
	configFile := ".jocasta"

	if len(os.Args) > 2 && os.Args[1] == "-c" {
		i = 3
		configPath = filepath.Dir(os.Args[2])
		configFile = filepath.Base(os.Args[2])
	}

	if (len(os.Args)-i) < 1 || os.Args[i] == "-h" {
		fmt.Println("usage: jocasta [-c configFileWithoutExtension] command to run with args")
		fmt.Println()
		fmt.Println("The config file name must be provided without the file extension, jocasta will try json, toml and yaml")
		os.Exit(0)
	}

	config, err := config.New(configPath, configFile, os.Args[i])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Will run %s\n", strings.Join(os.Args[i:], " "))
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

// will be called as goroutine to allow the stdout and stderr to be captured in parallel
func transfer(in io.ReadCloser, out io.Writer, wg *sync.WaitGroup) {
	if _, err := io.Copy(out, in); err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
