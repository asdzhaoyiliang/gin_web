package models

import "time"

type Category struct {
	Id      int
	Name    string
	Created time.Time
	Updated time.Time
}
