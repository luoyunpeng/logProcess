package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	logPath    string
	DataSource string

	readCh  chan string
	writeCh chan string
}

//read
func (logP *LogProcess) ReadLog() {
	line := "log info test"
	logP.readCh <- line
}

// log process
func (logP *LogProcess) Process() {
	logData := <-logP.readCh
	logP.writeCh <- strings.ToLower(logData)
}

// log write
func (logP *LogProcess) WriteTo(url string) {
	fmt.Println(<-logP.writeCh)
}

func main() {
	lp := &LogProcess{
		logPath:    "",
		DataSource: "",
		readCh:     make(chan string),
		writeCh:    make(chan string),
	}

	go lp.ReadLog()
	go lp.Process()
	go lp.WriteTo("")

	time.Sleep(time.Second * 1)
}
