package models

import "time"

type User struct {
	ID            int
	Email         string
	Username      string
	Password      string
	Token         string
	TokenDuration time.Time
}
