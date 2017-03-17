package view

import (
	"errors"

	"github.com/ieee0824/thor/util"
)

var View = map[util.LANG]*Views{}

type SingleView interface {
	Answer() string
	Message() string
	ToggleCursor()
	Controller()
	SetController(func())
}

type Views struct {
	key string
	now SingleView
	all map[string]SingleView
	fin bool
}

func New() *Views {
	return &Views{
		"",
		nil,
		map[string]SingleView{},
		false,
	}
}

func (v *Views) Add(view SingleView, k string) {
	v.all[k] = view
}

func (v *Views) Transition(k string) {
	if v.key != "" {
		v.all[v.key] = v.now
	}
	v.key = k
	v.now = v.all[k]
}

func (v *Views) GetView() (SingleView, error) {
	if v.now == nil {
		return nil, errors.New("view is not set")
	}
	return v.now, nil
}

func (v *Views) GetAnswer(key string) (string, error) {
	if view, ok := v.all[key]; ok {
		return view.Answer(), nil
	}
	return "", errors.New("can not find key " + key)
}

func (v *Views) Fin() {
	v.fin = true
}

func (v *Views) IsFin() bool {
	return v.fin
}
