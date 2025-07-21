package model

type ImageTask struct {
	ID       string
	Filename string
	Status   string // e.g. pending, processing, done
}
