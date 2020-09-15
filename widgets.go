package widgets

type Widget struct {
	id string
}

func (w Widget) ID() string {
	return w.id
}

