package models

import (
	"database/sql"
	"time"
)

type UserDTO struct {
	ID            *uint
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     *sql.NullTime
	Name          *string
	PassWord      *string
	Phone         *string
	Email         *string
	Identity      *string
	ClientIp      *string
	ClientPort    *string
	LoginTime     *time.Time
	HeartBeatTime *time.Time
	LoginOutTime  *time.Time
	IsLogout      *bool
	DeviceInfo    *string
}
