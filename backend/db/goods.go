package db

import (
	"time"

	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	GoodsId    uint   `gorm:"not null"`
	GoodsName  string `gorm:"not null"`
	AuthorName string `gorm:"not null"`
	Describe   string `gorm:"type:text;not null"`
	Images     string `gorm:"type:text"`
	Price      uint   `gorm:"not null"`
	Location   string `gorm:"type:text"`
}

type GoodsGet struct {
	GoodsId     uint
	GoodsName   string
	AuthorName  string
	Describe    string
	Images      string
	Price       uint
	Location    string
	CreatedTime time.Time
	UpdatedTime time.Time
	Avatar      string
}

type GoodsRequest struct {
	GoodsName  string
	AuthorName string
	Describe   string
	Images     string
	Price      uint
	Location   string
}

type PostCRUD struct{}
