package export

type export struct {
	Status string
	Error error
}

func (e *export) SetStatus(status string) {
	e.Status = status
}