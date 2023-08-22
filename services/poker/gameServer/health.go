package gameServer

import (
	"log"
	"os"
	"time"

	"agones.dev/agones/pkg/util/signals"
	sdk "agones.dev/agones/sdks/go"
)

func DoSignal() {
	signals.NewSigTermHandler(func() {
		log.Println("Exit signal received. Shutting down")
		os.Exit(0)
	})
}

func DoHealthPing(sdk *sdk.SDK, stop <-chan struct{}) {
	ticker := time.NewTicker(15 * time.Second)
	for {
		log.Print("Health Ping")
		err := sdk.Health()
		if err != nil {
			log.Fatalf("Could not send Health Ping")
		}
		select {
		case <-stop:
			ticker.Stop()
			log.Println("Stop health ping")
			return
		case <-ticker.C:
		}
	}
}
