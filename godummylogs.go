package godummylogs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type DummyLogs struct {
	Config Config

	sig      chan os.Signal
	stopChan chan bool
}

type LogTemplate struct {
	IntervalStr string   `json:"interval"`
	LogLines    []string `json:"logLines"`
}

func New(c string) *DummyLogs {
	_stopChan := make(chan bool)

	_sig := make(chan os.Signal, 1)
	signal.Notify(_sig, os.Interrupt, syscall.SIGINT, syscall.SIGHUP)

	dl := &DummyLogs{
		Config:   ParseConfig(c),
		stopChan: _stopChan,
		sig:      _sig,
	}

	go func() {
		_s := <-_sig
		switch _s {
		case os.Interrupt:
			dl.close()
		case syscall.SIGINT:
			dl.close()
		case syscall.SIGHUP:
			dl.close()
		}
	}()

	return dl
}

// generate
func (d DummyLogs) Generate() {
	// // auth.log
	// a := NewAuthLog(d.Config)
	// a.Generate()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}

func (d DummyLogs) close() {
	log.Println("Closing.....!")

	log.Println("Closed!")
	os.Exit(0)
}
