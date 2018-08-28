package do

import (
	"testing"
)

func TestH_MustJSON(t *testing.T) {

	str := string((H{
		"a": 1,
		"c": 2,
		"ddd": H{
			"aa": 1,
			"cc": 2,
		},
	}).MustJSON())

	if str != `{"a":1,"c":2,"ddd":{"aa":1,"cc":2}}` {
		t.Error("Result mismatch:" + str)
	}
}
