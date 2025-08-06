package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      ` + "`gorm:\"primaryKey\" json:\"id\"`" + `
	Name      string    ` + "`json:\"name\" gorm:\"type:varchar(100)\"`" + `
	Email     string    ` + "`json:\"email\" gorm:\"uniqueIndex;type:varchar(100)\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\" gorm:\"autoCreateTime\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\" gorm:\"autoUpdateTime\"`" + `
}
