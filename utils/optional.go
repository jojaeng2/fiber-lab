package utils

import "errors"

/*
Go에서 Optional은 권장되지 않음.
이유는 Go 언어의 철학 중 하나인 "오류를 처리하는 것은 호출자의 책임이다"와 관련이 있습니다.
*/
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
	return errors.New("null Pointer Exception")
}
