package updater

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sapphiregaze/discord-gorp/pkg/config"
	"github.com/sapphiregaze/discord-gorp/pkg/logger"
	"github.com/sapphiregaze/discord-gorp/pkg/rpc"
)

func Start() {
	cfg, err := config.Load()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to load config: %v", err))
		os.Exit(1)
	}

	client, err := rpc.NewClient()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to Discord: %v", err))
		os.Exit(1)
	}
	defer client.Close()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	activity := &cfg.Activity
	client.SetActivity(activity)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			client.SetActivity(activity)
		case <-sigs:
			logger.Info("Shutting down...")
			return
		}
	}
}
