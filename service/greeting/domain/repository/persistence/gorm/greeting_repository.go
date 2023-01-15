package gorm

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hb-chen/micro-starter/service/greeting/domain/model"
	"github.com/hb-chen/micro-starter/service/greeting/domain/repository"
)

type greetingRepository struct {
}

func NewGreetingRepository() repository.GreetingRepository {
	return &greetingRepository{}
}

func (r *greetingRepository) FindById(id int64) (*model.Msg, error) {
	item := model.Msg{}
	if result := db.Where("id = ?", id).First(&item); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	} else {
		return &item, nil
	}
}

func (r *greetingRepository) FindByMsg(msg string) (*model.Msg, error) {
	item := model.Msg{}
	if result := db.Where("msg = ?", msg).First(&item); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	} else {
		return &item, nil
	}
}

func (r *greetingRepository) Add(msg *model.Msg) error {
	err := db.Create(msg).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *greetingRepository) List(page, size int) ([]*model.Msg, error) {
	list := make([]*model.Msg, 0)
	err := db.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error

	return list, err
}
