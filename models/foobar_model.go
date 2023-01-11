package models

import "gorm.io/gorm"

type Foobar struct {
	gorm.Model
	What         string `json:"what"`
	Message      string `json:"message"`
	Registration uint64 `json:"registration"`
}
