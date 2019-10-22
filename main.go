package main

import (
	"fmt"
	"github.com/swallyu/gomqtt/broker"
	"log"
	"os"
	"os/signal"
)

func main() {

	server := broker.NewBroker()

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
	s := waitForExit()

	log.Panicln("exit for signal ", s)
}

func waitForExit() os.Signal {
	c := make(chan os.Signal)
	defer close(c)
	signal.Notify(c, os.Kill, os.Interrupt)
	s := <-c
	signal.Stop(c)
	fmt.Println(s)

	return s
}
