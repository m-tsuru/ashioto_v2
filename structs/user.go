package structs

import (
	"gorm.io/gorm"
	"time"
)

type UserBase struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey,autoIncrement" json:"id"`
	GoogleID  string    `gorm:"uniqueIndex" json:"google_id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	CreatedAt time.Time `json:"created_at" "gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserProfile struct {
	gorm.Model
	UserID      uint     `gorm:"uniqueIndex" json:"user_id"`
	Username    string   `gorm:"uniqueIndex" json:"username"` // @username
	ProfileName string   `json:"profile_name"`                // nickname
	AvatarURL   string   `json:"avatar_url"`                  // avatar (maybe google avatar)
	UserBase    UserBase `gorm:"foreignKey:UserID" json:"user"`
}

type UserExternalService struct {
	gorm.Model
	UserID              uint     `gorm:"uniqueIndex" json:"user_id"`
	SpotifyID           string   `gorm:"uniqueIndex" json:"spotify_id"`
	SpotifyToken        string   `json:"spotify_token"`
	SpotifyRefreshToken string   `json:"spotify_refresh_token"`
	UserBase            UserBase `gorm:"foreignKey:UserID" json:"user"`
}

type UserPublicProfile struct {
	Username    string `gorm:"uniqueIndex" json:"username"` // @username
	ProfileName string `json:"profile_name"`                // nickname
	AvatarURL   string `json:"avatar_url"`                  // avatar (maybe google avatar)
}
