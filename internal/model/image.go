package model

type ImageTask struct {
	ID       string `gorm:"primaryKey"`
	Filename string
	Status   string // e.g. pending, processing, done
}
