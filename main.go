package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
)

func main() {
	log.Printf("xwget\n")

	// selfUpdate := flag.NewFlagSet("self-update", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 1 && os.Args[1] == "self-update" {
		fmt.Println("Performing self update check")
		selfUpdate()
	}

	if flag.NFlag() == 0 && flag.NArg() == 0 {
		fmt.Println("Please run the command with -h for help.")
		return
	}

	// fmt.Println("os.Args: ", os.Args)
	// fmt.Println("flag.NFlag()", flag.NFlag())
	// fmt.Println("flag.NArg()", flag.NArg())

	args := os.Args
	rawURL, prefixedURL := "", ""
	mirror := os.Getenv("GITHUB_MIRROR")
	if mirror == "" {
		mirror = "https://ghproxy.com"
	}

	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "https://github.com") {
			rawURL = args[i]
			// args[i] = strings.Replace(args[i], "https://github.com", "https://ghproxy.com/https://github.com", -1)

			// mirrors.goproxyauth.com 只支持github.com这种格式的URL，即用来下载github release的URL
			prefixedURL = fmt.Sprintf("%s/%s", randomMirror(), rawURL)
		}

		if strings.Contains(args[i], "https://raw.githubusercontent.com") {
			rawURL = args[i]
			prefixedURL = fmt.Sprintf("%s/%s", mirror, rawURL)
		}

		if strings.Contains(args[i], "https://gist.github.com") {
			rawURL = args[i]
			prefixedURL = fmt.Sprintf("%s/%s", mirror, rawURL)
		}

		if strings.Contains(args[i], "https://gist.githubusercontent.com") {
			rawURL = args[i]
			prefixedURL = fmt.Sprintf("%s/%s", mirror, rawURL)
		}

		args[i] = prefixedURL
	}

	flag.Parse()

	// extract the file name from the URL path
	parsedURL, err := url.Parse(prefixedURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fileName := filepath.Base(parsedURL.Path)
	fmt.Println("File name:", fileName)

	// TOOD: extract filename from content-disposition header if filename is empty

	if flag.Parsed() && flag.Lookup("O") == nil {
		args = append(args, "-O")
		args = append(args, fileName)
	}

	execShell("wget", args[1:])
}

// https://go.dev/play/p/BV2GnfSiH1R
func execShell(cmd string, args []string) string {
	log.Println(cmd + " " + strings.Join(args, " "))
	var command = exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	var err = command.Start()
	if err != nil {
		return err.Error()
	}
	err = command.Wait()
	if err != nil {
		return err.Error()
	}
	return ""
}

func randomMirror() string {
	// rand.Seed(time.Now().UnixNano())

	// Generate a secure random seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Use the secure random seed to initialize the random number generator
	r.Seed(r.Int63())

	slice := []string{"https://ghproxy.com", "https://mirrors.goproxyauth.com"}
	// Generate a random index
	randomIndex := rand.Intn(len(slice))
	// Access the random element from the slice
	randomElement := slice[randomIndex]
	return randomElement
}

func selfUpdate() {
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/dfang/xwget",
			ArchiveName:   fmt.Sprintf("xwget_%s_%s.tar.gz", strings.Title(runtime.GOOS), "x86_64"),
		},
		ExecutableName: "xwget",
		// Version:        "v0.0.6", // You can change this value to trigger an update
		Version: xwgetVersion,
	}

	log.Println("Current version: " + u.Version)
	log.Println("Looking for updates...")
	var wg sync.WaitGroup
	wg.Add(1)
	// For the example we run the update in the background
	// but you could directly call u.Update()
	var updateStatus updater.UpdateStatus
	var updateErr error
	go func() {
		updateStatus, updateErr = u.Update()
		wg.Done()
	}()

	wg.Wait() // Waiting for the update process to finish before exiting
	if updateErr != nil {
		log.Println(updateErr)
	}
	if updateStatus == updater.Updated {
		log.Println("Updated!")
	}
}
