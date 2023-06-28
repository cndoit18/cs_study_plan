package chainofresponsibility

type Compoent interface {
	Handles(any) bool
}

type Handler interface {
	Compoent
	SetNext(Compoent) Handler
}

type CompoentFunc func(any) bool

func (c CompoentFunc) Handles(i any) bool {
	return c(i)
}

func (c CompoentFunc) SetNext(cm Compoent) Handler {
	return CompoentFunc(func(a any) bool {
		if !c(a) {
			return false
		}
		return cm.Handles(a)
	})
}
