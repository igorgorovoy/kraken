package main

import (
	"code.uber.internal/go-common.git/x/log"

	"code.uber.internal/infra/kraken/lib/dockerregistry"
	"code.uber.internal/infra/kraken/lib/serverset"
	"code.uber.internal/infra/kraken/lib/store"
	"code.uber.internal/infra/kraken/lib/torrent"
	"code.uber.internal/infra/kraken/metrics"
)

// Config defines agent configuration.
type Config struct {
	Logging  log.Configuration     `yaml:"logging"`
	Metrics  metrics.Config        `yaml:"metrics"`
	Store    store.Config          `yaml:"store"`
	Registry dockerregistry.Config `yaml:"registry"`
	Torrent  torrent.Config        `yaml:"torrent"`
	Tracker  TrackerConfig         `yaml:"tracker"`
}

// TrackerConfig defines configuration for agent's dependency on tracker.
type TrackerConfig struct {
	RoundRobin serverset.RoundRobinConfig `yaml:"round_robin"`
}