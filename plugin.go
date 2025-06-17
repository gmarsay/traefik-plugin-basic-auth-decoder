package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}

type BasicAuthDecoder struct {
	next   http.Handler
	config *Config
}

func New(next http.Handler, config *Config, name string) (http.Handler, error) {
	return &BasicAuthDecoder{
		next:   next,
		config: config,
	}, nil
}

func (p *BasicAuthDecoder) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")

	if authHeader != "" && strings.HasPrefix(authHeader, "Basic ") {
		encoded := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err == nil {
			parts := strings.SplitN(string(decoded), ":", 2)
			if len(parts) == 2 {
				req.Header.Set("X-Traefik-Loggable-Username", parts[0])
			}
		}
	}

	p.next.ServeHTTP(rw, req)
} 