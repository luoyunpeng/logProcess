package main

import (
	"fmt"
	"strings"
	"time"
)

type LogReader interface {
	Read(ch chan string)
}

type LogRead struct {
	path string
}

func (read *LogRead) Read(ch chan string) {
	ch <- "log test"
}

type LogWriter interface {
	Write(ch chan string)
}

type LogWrite struct {
	DataSource string
}

func (write *LogWrite) Write(ch chan string) {
	fmt.Println(<-ch)
}

type LogProcess struct {
	readCh  chan string
	writeCh chan string

	LogReader
	LogWriter
}

// log process
func (logP *LogProcess) Process() {
	logData := <-logP.readCh
	logP.writeCh <- strings.ToLower(logData)
}

func main() {
	lp := &LogProcess{
		readCh:    make(chan string),
		writeCh:   make(chan string),
		LogReader: &LogRead{},
		LogWriter: &LogWrite{},
	}

	go lp.Read(lp.readCh)
	go lp.Process()
	go lp.Write(lp.writeCh)

	time.Sleep(time.Second * 1)
}
