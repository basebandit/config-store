package configstore

import "gorm.io/gorm"

type KV struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex;not null"`
	Value string `gorm:"not null"`
}
