package model

import "time"

type ImageTask struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	Filename string    `json:"filename"`
	Type     string    `json:"type"`   // e.g. resize, grayscale, blur
	Status   string    `json:"status"` // pending, processing, done, failed
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
