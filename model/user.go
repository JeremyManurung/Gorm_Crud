package model

import ()

type User struct {
	Id	      uint      `gorm:"primarykey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}
