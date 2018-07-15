package math

import (
	"encoding/json"

	simplejson "github.com/bitly/go-simplejson"
)

type Math struct {
	A int
	B int
}

func (m *Math) Divide() int {
	return m.A / m.B
}

func (m *Math) GetA() int {
	return m.A
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
