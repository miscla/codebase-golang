package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	AppPort string
	GinMode string

	CacheTTLSeconds int
}

func Load() *Config {
	// you can use github.com/joho/godotenv to load .env in development,
	// but here we'll assume env already set (or you can enable godotenv).
	if os.Getenv("DB_HOST") == "" {
		// try load .env file silently (optional)
		_ = loadDotEnv()
	}

	cacheTTL := 30
	if v := os.Getenv("CACHE_TTL"); v != "" {
		if t, err := strconv.Atoi(v); err == nil {
			cacheTTL = t
		}
	}

	cfg := &Config{
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres"),
		DBName:          getEnv("DB_NAME", "yourdb"),
		DBSSLMode:       getEnv("DB_SSLMODE", "disable"),
		AppPort:         getEnv("APP_PORT", "8080"),
		GinMode:         getEnv("GIN_MODE", "release"),
		CacheTTLSeconds: cacheTTL,
	}

	log.Printf("Loaded config: host=%s port=%s db=%s", cfg.DBHost, cfg.DBPort, cfg.DBName)
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// optional small helper to load .env using os.ReadFile if present
func loadDotEnv() error {
	// keep this minimal to avoid extra dependency; parse KEY=VALUE lines
	data, err := os.ReadFile(".env")
	if err != nil {
		return err
	}
	lines := string(data)
	for _, line := range splitLines(lines) {
		if line == "" || line[0] == '#' {
			continue
		}
		// simple split at first '='
		for i := 0; i < len(line); i++ {
			if line[i] == '=' {
				key := line[:i]
				val := line[i+1:]
				_ = os.Setenv(key, val)
				break
			}
		}
	}
	return nil
}

func splitLines(s string) []string {
	out := []string{}
	cur := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '\r' {
			continue
		}
		if c == '\n' {
			out = append(out, cur)
			cur = ""
			continue
		}
		cur += string(c)
	}
	if cur != "" {
		out = append(out, cur)
	}
	return out
}
