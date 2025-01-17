package configstore

import (
	"errors"

	"gorm.io/gorm"
)

type KVService struct {
	db *gorm.DB
}

func NewKVService(db *gorm.DB) *KVService {
	return &KVService{db}
}

func (s *KVService) ListKVS() ([]KV, error) {
	var kvs []KV
	if result := s.db.Find(&kvs); result.Error != nil {
		return nil, result.Error
	}
	return kvs, nil
}

func (s *KVService) GetKV(key string) (*KV, error) {
	var kv KV
	if result := s.db.Where("key = ?", key).First(&kv); result.Error != nil {
		return nil, result.Error
	}
	return &kv, nil
}

func (s *KVService) CreateKV(key, value string) (*KV, error) {
	var existingKV KV
	if result := s.db.Where("key = ?", key).First(&existingKV); result.Error == nil {
		return nil, errors.New("key already present in DB")
	}
	kv := KV{
		Key:   key,
		Value: value,
	}
	if result := s.db.Create(&kv); result.Error != nil {
		return nil, result.Error
	}
	return &kv, nil
}

func (s *KVService) UpdateKV(key, value string) (*KV, error) {
	var kv KV
	if result := s.db.Where("key = ?", key).First(&kv); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("key not found")
		}
		return nil, result.Error
	}
	kv.Value = value
	if result := s.db.Save(&kv); result.Error != nil {
		return nil, result.Error
	}
	return &kv, nil
}

func (s *KVService) DeleteKV(key string) error {
	if result := s.db.Where("key = ?", key).Delete(&KV{}); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("key not found")
	}
	return nil
}
