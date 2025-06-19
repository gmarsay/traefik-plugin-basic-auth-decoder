package basicAuthDecoder

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
)

const defaultHeader = "X-Traefik-Loggable-Username"

type Config struct {
	HeaderName string
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader != "" {
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) == 2 && authHeaderParts[0] == "Basic" {
				decoded, err := base64.StdEncoding.DecodeString(authHeaderParts[1])
				if err == nil {
					r.Header.Add(config.HeaderName, string(decoded))
				}
			}
		}

		next.ServeHTTP(rw, r)
	}), nil
}
