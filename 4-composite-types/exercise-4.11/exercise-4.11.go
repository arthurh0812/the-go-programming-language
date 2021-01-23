package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var wd string

func exit() {
	fmt.Println("exiting...")
}

func clear() {
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(wd)
}

func login() {
	fmt.Println("logging in...")

}

func logout() {
	fmt.Println("loging out...")
}

func main() {
	wd, _ = os.Getwd()

	fmt.Println(wd)

	// CLI commands
	commands := map[string]func(){
		"exit":   exit,
		"login":  login,
		"logout": logout,
		"create": createIssue,
		"read":   readIssue,
		"update": updateIssue,
		"close":  closeIssue,
		"clear":  clear,
	}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		cmd, ok := commands[s.Text()]
		if !ok {
			log.Println("Access failed. Command does not exist.")
			continue
		}

		// execute the command
		cmd()
	}
}

func createIssue() {
	fmt.Println("creating...")
}

func readIssue() {
	fmt.Println("reading...")
}

func updateIssue() {
	fmt.Println("updating...")
}

func closeIssue() {
	fmt.Println("closing...")
}
