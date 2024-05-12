package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	DefaultAddr        = "localhost:8080"
	DefaultServerURL   = "http://localhost:8080"
	DefaultRoutePrefix = "/"
)

var (
	Addr            string
	ServerURLPrefix string
	RoutePrefix     string
)

func ParseFlags() error {
	flag.StringVar(&Addr, "a", DefaultAddr, "address and port to run server")
	flag.StringVar(&ServerURLPrefix, "b", DefaultServerURL, "server base url prefix to use for requests")

	flag.Parse()

	if envSrvAddr, found := os.LookupEnv("SERVER_ADDRESS"); found {
		Addr = envSrvAddr
	}

	if envBaseURL, found := os.LookupEnv("BASE_URL"); found {
		ServerURLPrefix = envBaseURL
	}

	ServerURLPrefix = strings.TrimSuffix(ServerURLPrefix, "/")

	u, err := url.Parse(ServerURLPrefix)
	if err != nil {
		return fmt.Errorf("server base url parse error: %w", err)
	}

	RoutePrefix = strings.TrimSuffix(u.Path, "/")
	if RoutePrefix == "" {
		RoutePrefix = DefaultRoutePrefix
	}

	return nil
}
