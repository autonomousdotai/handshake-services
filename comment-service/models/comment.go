package models

import (
	_ "encoding/gob"
	"time"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Comment struct {
	DateCreated  time.Time
	DateModified time.Time
	ID           int64
	UserId       int64
	ObjectId     string
	Comment      string
	Image        string
	Address      string
	Status       int
	User         User
}

func (Comment) TableName() string {
	return "comment"
}
