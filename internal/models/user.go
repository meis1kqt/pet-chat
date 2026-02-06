package models

import "modernc.org/libc/uuid"


type User struct {
	ID uuid
	Email string
	PassHash []byte
}