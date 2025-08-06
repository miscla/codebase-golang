package service

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	"codebase-golang/pkg/config"
	"codebase-golang/pkg/logger"
)

// User model
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	db       *sql.DB
	cache    []User
	mu       sync.RWMutex
	lastSync time.Time
	ttl      time.Duration
	cfg      *config.Config
}

func NewUserService(db *sql.DB, cfg *config.Config) *UserService {
	return &UserService{
		db:  db,
		ttl: time.Duration(cfg.CacheTTLSeconds) * time.Second,
		cfg: cfg,
	}
}

// Fetch from DB and update cache (force)
func (s *UserService) RefreshCache() error {
	if s.db == nil {
		return errors.New("db is nil")
	}
	rows, err := s.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return err
		}
		users = append(users, u)
	}
	s.mu.Lock()
	s.cache = users
	s.lastSync = time.Now()
	s.mu.Unlock()

	logger.Info("user cache refreshed", len(users))
	return nil
}

// GetAll returns cached users if TTL not expired; otherwise refresh and return.
func (s *UserService) GetAll() ([]User, error) {
	s.mu.RLock()
	if time.Since(s.lastSync) < s.ttl && len(s.cache) > 0 {
		users := make([]User, len(s.cache))
		copy(users, s.cache)
		s.mu.RUnlock()
		return users, nil
	}
	s.mu.RUnlock()

	// refresh synchronously
	if err := s.RefreshCache(); err != nil {
		return nil, err
	}

	s.mu.RLock()
	users := make([]User, len(s.cache))
	copy(users, s.cache)
	s.mu.RUnlock()
	return users, nil
}
