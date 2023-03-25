package optional

type Optional struct {
	instance interface{}
}

func OfNullable(instance interface{}) *Optional {
	if instance == nil {
		return &Optional{instance: nil}
	}
	return &Optional{instance: instance}
}

func (optional Optional) IsPresent() bool {
	return optional.instance != nil
}

func (optional Optional) OrElse(instance interface{}) interface{} {
	if optional.IsPresent() {
		return optional.instance
	}
	return instance
}

func (optional Optional) OrElseThrow() interface{} {
	if optional.IsPresent() {
		return optional.instance
	}
	panic("Null Pointer Exception")
}
