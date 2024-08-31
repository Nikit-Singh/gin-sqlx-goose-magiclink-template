package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Token     uuid.UUID `db:"token"`
	UserID    uuid.UUID `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
}
