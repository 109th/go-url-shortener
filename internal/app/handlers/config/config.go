package config

import (
	"flag"
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

func ParseFlags() {
	flag.StringVar(&Addr, "a", DefaultAddr, "address and port to run server")
	flag.StringVar(&ServerURLPrefix, "b", DefaultServerURL, "server base url prefix to use for requests")

	flag.Parse()

	if envSrvAddr := os.Getenv("SERVER_ADDRESS"); envSrvAddr != "" {
		Addr = envSrvAddr
	}

	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		ServerURLPrefix = envBaseURL
	}

	ServerURLPrefix = strings.TrimSuffix(ServerURLPrefix, "/")

	u, err := url.Parse(ServerURLPrefix)
	if err != nil {
		panic(err)
	}

	RoutePrefix = strings.TrimSuffix(u.Path, "/")
	if RoutePrefix == "" {
		RoutePrefix = DefaultRoutePrefix
	}
}
