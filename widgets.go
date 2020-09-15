package widgets

import uuid "github.com/satori/go.uuid"

type widget struct {
	id string
}

func NewWidget() widget {
	return widget{
		id: uuid.NewV4().String(),
	}
}

func (w widget) ID() string {
	return w.id
}
