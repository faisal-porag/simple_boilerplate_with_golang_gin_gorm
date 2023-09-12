package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"golang_boilerplate_with_gin/config"
	"golang_boilerplate_with_gin/httpserver"
	"golang_boilerplate_with_gin/state"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Config parsing failed")
	}

	appState := state.NewState(cfg)
	httpserver.Server(appState)

}
