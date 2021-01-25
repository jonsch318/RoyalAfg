package gameServer

import (
	"log"
	"os"
	"time"

	"agones.dev/agones/pkg/util/signals"
	sdk "agones.dev/agones/sdks/go"
)

func DoSignal() {
	stop := signals.NewStopChannel()
	<-stop
	log.Println("Exit signal received. Shutting down")
	os.Exit(0)
}

func DoHealthPing(sdk *sdk.SDK, stop <-chan struct{}) {
	tick := time.Tick(15 * time.Second)
	for {
		log.Print("Health Ping")
		err := sdk.Health()
		if err != nil {
			log.Fatalf("Could not send Health Ping")
		}
		select {
		case <-stop:
			log.Println("Stop health ping")
			return
		case <-tick:
		}
	}
}
