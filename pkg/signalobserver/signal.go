package signalobserver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// ObserveSignal observes signals from os
func ObserveSignal(callback func()) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	fmt.Printf("Received signalobserver: %s \n", sig)
	callback()
}
