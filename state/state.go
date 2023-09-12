package state

import (
	"github.com/rs/zerolog/log"
	"golang_boilerplate_with_gin/config"
	"golang_boilerplate_with_gin/db"
)

type State struct {
	Cfg        *config.Config
	Repository *db.PgRepository
}

func NewState(cfg *config.Config) *State {
	pgDb, err := db.NewPgRepository(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("pg connection error")
	} else {
		log.Info().Msg("db connection successful")
	}

	return &State{
		Cfg:        cfg,
		Repository: pgDb,
	}
}
