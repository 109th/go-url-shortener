package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	defaultAddr        = "localhost:8080"
	defaultServerURL   = "http://localhost:8080"
	defaultRoutePrefix = "/"
)

type Config struct {
	Addr            string
	ServerURLPrefix string
	RoutePrefix     string
}

func ParseFlags() (*Config, error) {
	cfg := &Config{
		Addr:            defaultAddr,
		ServerURLPrefix: defaultServerURL,
		RoutePrefix:     defaultRoutePrefix,
	}

	flag.StringVar(&cfg.Addr, "a", defaultAddr, "address and port to run server")
	flag.StringVar(&cfg.ServerURLPrefix, "b", defaultServerURL, "server base url prefix to use for requests")

	flag.Parse()

	if envSrvAddr, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		cfg.Addr = envSrvAddr
	}

	if envBaseURL, ok := os.LookupEnv("BASE_URL"); ok {
		cfg.ServerURLPrefix = envBaseURL
	}

	cfg.ServerURLPrefix = strings.TrimSuffix(cfg.ServerURLPrefix, "/")

	u, err := url.Parse(cfg.ServerURLPrefix)
	if err != nil {
		return nil, fmt.Errorf("server base url parse error: %w", err)
	}

	cfg.RoutePrefix = strings.TrimSuffix(u.Path, "/")
	if cfg.RoutePrefix == "" {
		cfg.RoutePrefix = defaultRoutePrefix
	}

	return cfg, nil
}
