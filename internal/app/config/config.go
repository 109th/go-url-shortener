package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	defaultAddr            = "localhost:8080"
	defaultServerURL       = "http://localhost:8080"
	defaultRoutePrefix     = "/"
	defaultFileStoragePath = "/tmp/short-url-db.json"
	defaultStorageType     = "file"
)

type Config struct {
	Addr            string
	ServerURLPrefix string
	RoutePrefix     string
	FileStoragePath string
	StorageType     string
}

func ParseFlags() (*Config, error) {
	cfg := &Config{
		Addr:            defaultAddr,
		ServerURLPrefix: defaultServerURL,
		RoutePrefix:     defaultRoutePrefix,
		StorageType:     defaultStorageType,
	}

	flag.StringVar(&cfg.Addr, "a", defaultAddr, "address and port to run server")
	flag.StringVar(&cfg.ServerURLPrefix, "b", defaultServerURL, "server base url prefix to use for requests")
	flag.StringVar(&cfg.FileStoragePath, "f", defaultFileStoragePath, "path to the data store file")
	flag.StringVar(&cfg.StorageType, "s", defaultStorageType, "storage type, possible values: memory, file")

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

	if envFileStoragePath, ok := os.LookupEnv("FILE_STORAGE_PATH"); ok {
		cfg.FileStoragePath = envFileStoragePath
	}

	if envStorageType, ok := os.LookupEnv("STORAGE_TYPE"); ok {
		cfg.StorageType = envStorageType
	}

	return cfg, nil
}
