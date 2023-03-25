package optional

type Optional interface {
	IsPresent() bool

	OrElse() interface{}

	OrElseThrow(interface{}) interface{}
}

type OptionalImpl struct {
	instance interface{}
}

func (optionalImpl OptionalImpl) IsPresent() bool {
	return optionalImpl.instance != nil
}

func (optionalImpl OptionalImpl) OrElse(instance interface{}) interface{} {
	if optionalImpl.IsPresent() {
		return optionalImpl.instance
	}
	return instance
}

func (optionalImpl OptionalImpl) OrElseThrow() interface{} {
	if optionalImpl.IsPresent() {
		return optionalImpl.instance
	}
	panic("Null Pointer Exception")
}
