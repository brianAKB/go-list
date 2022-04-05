package list

import (
	"fmt"
	"reflect"
)

type List interface {
	Append(interface{}) error
	Remove(uint) error
	Clear()
	Length() int
	At(uint) (interface{}, error)
}

type list struct {
	arr         []interface{}
	myType      reflect.Type
	initialized bool
}

func (l *list) Append(v interface{}) error {
	if !l.initialized {
		l.arr = append(l.arr, v)
		l.myType = reflect.TypeOf(v)
		l.initialized = true
		return nil
	} else if t := reflect.TypeOf(v); t == l.myType {
		l.arr = append(l.arr, v)
		return nil
	}
	return fmt.Errorf("value %v is incorrect type: Should be type %s", v, l.myType.Name())
}

func (l *list) At(i uint) (interface{}, error) {
	if i >= uint(len(l.arr)) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	return l.arr[i], nil
}

func (l *list) Remove(i uint) error {
	if i >= uint(len(l.arr)) {
		return fmt.Errorf("index %d out of bounds", i)
	}
	l.arr = append(l.arr[:i], l.arr[i+1:]...)
	return nil
}

func (l *list) Clear() {
	l.arr = make([]interface{}, 0)
}

func (l *list) Length() int {
	return len(l.arr)
}

func NewList() List {
	l := list{initialized: false, arr: make([]interface{}, 0)}
	return &l
}
