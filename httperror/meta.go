package httperror

type Meta struct {
	Code int
	Data interface{}
}

func (m Meta) SetData(data interface{}) Meta {
	m.Data = data
	return m
}

func NewMeta(code int) Meta {
	return Meta{Code: code}
}
