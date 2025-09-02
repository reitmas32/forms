package main

import (
	"common/domain/logger"
	"fomrs/internal/core/server"
	"fomrs/internal/core/settings"
	"log"
)

func main() {
	settings.LoadDotEnv()

	settings.LoadEnvs()

	logger.InitLogger(settings.Settings.ENVIRONMENT, "fomrs", settings.Settings.LOKI_URL)

	switch settings.Settings.DEPLOY_MODE {
	case settings.DeployModeAPI:
		server.Run()
	case settings.DeployModeLambda:
		server.RunLambda()
	default:
		log.Fatalf("Invalid deploy mode: %s", settings.Settings.DEPLOY_MODE)
	}
}
