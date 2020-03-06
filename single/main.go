package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGUSR1)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT:
			log.Println("SIGSTOP")
			return
		case syscall.SIGSTOP:
			log.Println("SIGSTOP")
			return
		case syscall.SIGHUP:
			log.Println("SIGHUP")
			return
		case syscall.SIGKILL:
			log.Println("SIGKILL")
			return
		case syscall.SIGUSR1:
			log.Println("SIGUSR1")
			return
		default:
			log.Println("default")
			return
		}
	}
}
