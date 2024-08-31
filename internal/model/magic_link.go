package model

import "time"

type MagicLink struct {
	ID        int64     `db:"id"`
	Email     string    `db:"email"`
	OTP       string    `db:"otp"`
	IsUsed    bool      `db:"is_used"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}
