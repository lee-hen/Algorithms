// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This example is provided with help by Gabriel Aszalos.

// This sample program demonstrates how to create a simple chat system.
package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/lee-hen/Algorithms/5_context_or_beyond/101_go_routines/17_concurrency_patterns/chat"
)

func main() {
	cr := chat.New()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting Down Started")
	cr.Close()
	log.Println("Shutting Down Completed")
}
