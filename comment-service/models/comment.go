package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm"
	_ "encoding/gob"
	"time"
)

type Comment struct {
	DateCreated  time.Time
	DateModified time.Time
	ID           int64
	UserId       int64
	ObjectType   string
	ObjectId     string
	Comment      string
	Image        string
	Status       int
	User         User
}

func (Comment) TableName() string {
	return "comment"
}
