package main

import (
	"context"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Host string `json:"host"`
	SQL  string `json:"sql"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
