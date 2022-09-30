package models

import "time"

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	LoginCount int
	LastTime   time.Time
	LastIp     int
	State      int8
	Created    time.Time
	Updated    time.Time
}
