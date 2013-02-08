package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
)

func main() {

	fmt.Println("thru - The worlds simplest web server")
	fmt.Println("by Mat Ryer and Tyler Bunnell")
	fmt.Println("")

	// get the address from the first argument
	var address string
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	// set default
	if len(address) < 3 {
		address = "localhost:4000"
	}

	// add 'localhost' if they just specified a port
	if strings.HasPrefix(address, ":") {
		address = fmt.Sprintf("localhost%s", address)
	}

	s := &http.Server{
		Addr:           address,
		Handler:        http.FileServer(http.Dir("./")),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", address)

	if listenErr != nil {
		fmt.Printf("FAILED: :-( Could not listen: %s\n", listenErr)
		os.Exit(2)
		return
	}

	go func() {
		for _ = range c {

			// sig is a ^C, handle it

			// stop the HTTP server
			fmt.Print("\n\nStopping the server...\n")
			listener.Close()

			fmt.Print("Finished - bye bye.  ;-)\n")

		}
	}()

	fmt.Printf("These files are going thru %s\n", address)
	fmt.Print("\nUse Ctrl+C to stop\n")

	// should we open their browser for them the lazy bastards?
	var openBrowser bool
	if len(os.Args) > 1 && os.Args[1] == "-o" {
		openBrowser = true
	} else if len(os.Args) > 2 && os.Args[2] == "-o" {
		openBrowser = true
	}

	if openBrowser {

		command := exec.Command("open", fmt.Sprintf("http://%s", address))
		command.Run()

	}

	// begin the server
	fmt.Printf("Fatal error: %s\n", s.Serve(listener))

}
