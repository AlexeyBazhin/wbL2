package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"

	ps "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	curPath := "dev8shell:/$ "
	for {
		fmt.Print(curPath)
		// input := bufio.NewScanner(os.Stdin)
		// input.Scan()
		// command := strings.Split(input.Text(), " ")

		command := []string{"exec", "telegram-desktop"}
		length := len(command)
		switch command[0] {
		case "cd":
			if length == 1 {
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				err = os.Chdir(home)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			} else {
				err := os.Chdir(command[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			}
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			curPath = path.Join("dev8shell:/", pwd, "$ ")
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(pwd)
		case "echo":
			fmt.Println(strings.Join(command[1:], " "))
		case "kill":
			if length == 1 {
				fmt.Fprintln(os.Stderr, "kill: not enough arguments")
				continue
			}
			pid, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			proc, err := os.FindProcess(pid)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			err = proc.Kill()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
		case "ps":
			w := new(tabwriter.Writer)
			w = w.Init(os.Stdout, 5, 30, 5, ' ', 0)
			list, err := ps.Processes()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Fprintln(w, "NAME\tPID\tPPID")
			for _, p := range list {
				fmt.Fprintf(w, "%s\t%d\t%d\n", p.Executable(), p.Pid(), p.PPid())
			}
			w.Flush()
		case "fork":
			syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		case "exec":
			cmd := exec.Command(command[1], command[2:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "exit":
			os.Exit(0)
		}

	}
}
