package dao

import (
	"../models"
	"log"
	"github.com/jinzhu/gorm"
	"time"
)

type CommentDao struct {
}

func (commentDao CommentDao) GetById(id int64) (models.Comment) {
	dto := models.Comment{}
	err := models.Database().Where("id = ?", id).First(&dto).Error
	if err != nil {
		log.Print(err)
	}
	return dto
}

func (commentDao CommentDao) Create(dto models.Comment, tx *gorm.DB) (models.Comment, error) {
	if tx == nil {
		tx = models.Database()
	}
	dto.DateCreated = time.Now()
	dto.DateModified = dto.DateCreated
	err := tx.Create(&dto).Error
	if err != nil {
		log.Println(err)
		return dto, err
	}
	return dto, nil
}

func (commentDao CommentDao) Update(dto models.Comment, tx *gorm.DB) (models.Comment, error) {
	if tx == nil {
		tx = models.Database()
	}
	dto.DateModified = time.Now()
	err := tx.Save(&dto).Error
	if err != nil {
		log.Println(err)
		return dto, err
	}
	return dto, nil
}

func (commentDao CommentDao) Delete(dto models.Comment, tx *gorm.DB) (models.Comment, error) {
	if tx == nil {
		tx = models.Database()
	}
	err := tx.Delete(&dto).Error
	if err != nil {
		log.Println(err)
		return dto, err
	}
	return dto, nil
}
