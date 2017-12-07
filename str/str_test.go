package str

import "testing"

func TestFormat(t *testing.T) {
	str := "http://baidu.com/{path}/{name}/{name}.jpg"
	params := map[string]interface{}{
		"path": "aaaa",
		"name": "this-is-name",
	}

	t.Log(Format(str, params))
}
