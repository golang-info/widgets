package widgets

import uuid "github.com/satori/go.uuid"

// Widget is a ...
type Widget interface {
	// ID 返回这个 widget 的唯一标识符
	ID() string
}

type widget struct {
	id string
}

func NewWidget() Widget {
	return widget{
		id: uuid.NewV4().String(),
	}
}

func (w widget) ID() string {
	return w.id
}
