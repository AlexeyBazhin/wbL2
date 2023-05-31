package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type TelnetOptions struct {
	timeoutStr string // -t
}

func ParseFlags() *TelnetOptions {
	options := &TelnetOptions{}

	// Определение флагов командной строки
	flag.StringVar(&options.timeoutStr, "t", "10s", "указать timeout")

	// Парсинг аргументов командной строки
	flag.Parse()

	return options
}

func main() {
	options := ParseFlags()

	timeout, err := time.ParseDuration(options.timeoutStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse timeout: ", err)
		os.Exit(1)
	}

	if err := Telnet("opennet.ru", "80", timeout); err != nil {
		fmt.Fprintln(os.Stderr, "telnet error: ", err)
		os.Exit(1)
	}
}

func Telnet(host, port string, timeout time.Duration) error {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	errChan := make(chan error, 1)
	go send(conn, signalChan, errChan)
	go recieve(conn, errChan)

	select {
	case <-signalChan:
		conn.Close()
	case err = <-errChan:
		if err != nil {
			return err
		}
	}
	return nil
}

func send(conn net.Conn, signalChan chan<- os.Signal, errChan chan<- error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		// if _, err := conn.Write([]byte(msg)); err != nil {

		// 	errChan <- err
		// }
		
		fmt.Fprintln(conn, msg)
	}

	if scanner.Err() != nil {
		errChan <- scanner.Err()
		return
	}
	signalChan <- syscall.Signal(syscall.SIGQUIT)
}

func recieve(conn net.Conn, errChan chan<- error) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		errChan <- scanner.Err()
	}
	errChan <- fmt.Errorf("connection closed by host")
}
