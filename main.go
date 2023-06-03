package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func execShell(cmd string, args []string) string {
	var argss = ""
	for i := 0; i < len(args); i++ {
		argss = argss + args[i] + " "
	}
	log.Println(cmd + " " + string(argss))
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

func main() {
	log.Printf("xwget\n")
	args := os.Args
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "https://github.com") {
			args[i] = strings.Replace(args[i], "https://github.com", "https://ghproxy.com/https://github.com", -1)
		}
		
		if strings.Contains(args[i], "https://raw.githubusercontent.com") {
			args[i] = strings.Replace(args[i], "https://raw.githubusercontent.com", "https://ghproxy.com/https://raw.githubusercontent.com", -1)
		}

		if strings.Contains(args[i], "https://gist.github.com") {
			args[i] = strings.Replace(args[i], "https://gist.github.com", "https://ghproxy.com/https://gist.github.com", -1)
		}

		if strings.Contains(args[i], "https://gist.githubusercontent.com") {
			args[i] = strings.Replace(args[i], "https://gist.githubusercontent.com", "https://ghproxy.com/https://gist.githubusercontent.com", -1)
		}
	}
	execShell("wget", args[1:])
}
