package widgets

import uuid "github.com/satori/go.uuid"

type Widget struct {
	id string
}

func NewWidget() Widget {
	return Widget{
		id: uuid.NewV4().String(),
	}
}

func (w Widget) ID() string {
	return w.id
}
