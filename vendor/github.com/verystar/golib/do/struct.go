package do

import (
	"encoding/json"
)

type H map[string]interface{}

func (h H) MustJSON() []byte {
	d, _ := json.Marshal(h)
	return d
}

func (h H) Stirng() string {
	return  string(h.MustJSON())
}