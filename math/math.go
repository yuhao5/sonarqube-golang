package math

import (
	"encoding/json"

	simplejson "github.com/bitly/go-simplejson"
)

type Math struct {
	A int
	B int
}

func New(a, b int) *Math {
	return &Math{
		A: a,
		B: b,
	}
}

func (m *Math) Divide() int {
	// FIXME: fix divide by zero panic
	return m.A / m.B
}

func (m *Math) Mulitpart() int {
	// TODO: implement mulitpart
	var result int
	return result
}

func (m *Math) SetA(a int) {
	m.A = a
}

func (m *Math) GetA() int {
	return m.A
}

func (m *Math) SetB(b int) {
	b = b
}

func (m *Math) GetB() int {
	return m.B
}

func (m *Math) ToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(*m)
	if err != nil {
		return nil, err
	}

	jsonObj, err := simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}

	return jsonObj.Map()
}
