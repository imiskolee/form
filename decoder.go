package form

import (
	"fmt"
	"net/url"
	"strings"
)

/**
用于支持PHP的$_POST解析。

```go

	query := `a[b][c]=1&a[b][c]=1&a[b][d]=1&f=1&g[h][i]=1&g[h][i][t]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()   //desp 是 map[string]interface{}

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)

```

**/
type Form struct {
	dest        map[string]interface{}
	raw         string
	state       int
	currentPath []string
	current     string
	needQueryUnescape bool
}

func NewForm(raw string) *Form {
	form := new(Form)
	form.raw = raw
	form.NeedQueryUnescape(true)
	return form
}

func (form *Form)NeedQueryUnescape(need bool) {
	form.needQueryUnescape = need
}

func (form *Form) reset() {
	form.dest = make(map[string]interface{})
}

func (form *Form) Decode() (map[string]interface{}, error) {

	form.reset()

	u := form.raw

	vals := make(map[string]interface{})

	current := ""

	for _, c := range u {

		switch c {

		case '&':
			pair := current
			index := strings.Index(pair, "=")
			if index > 0 {
				key := pair[:index]
				val := pair[index+1:]

				insertValue(&vals, key, val,form.needQueryUnescape)
			} else {
				insertValue(&vals, pair, "",form.needQueryUnescape)
			}
			current = ""
			continue
		default:
			current = current + string(c)

		}
	}

	if len(current) > 0 {
		pair := current
		index := strings.Index(pair, "=")
		if index > 0 {
			key := pair[:index]
			val := pair[index+1:]
			insertValue(&vals, key, val,form.needQueryUnescape)
		} else {
			insertValue(&vals, pair, "",form.needQueryUnescape)
		}
		current = ""
	}
	return form.parseArray(vals), nil
}

func insertValue(destP *map[string]interface{}, key string, val string,need bool) {
	key,_ = url.PathUnescape(key)

	var path []string
	var current string

	for _, c := range key {
		switch c {
		case '[':
			if len(current) > 0 {
				path = append(path, current)
				current = ""
			}
			break
		case ']':
			path = append(path, current)
			current = ""
			continue
		default:
			current += string(c)
		}
	}

	if len(current) > 0 {
		path = append(path, current)
	}

	dest := *destP
	for i := 0; i < len(path)-1; i++ {
		p := path[i]
		if p == "" {
			p = fmt.Sprint(len(dest))
		}
		if _, ok := dest[p].(map[string]interface{}); !ok {
			dest[p] = make(map[string]interface{})
		}
		dest = dest[p].(map[string]interface{})
	}
	p := path[len(path)-1]
	if p == "" {
		p = fmt.Sprint(len(dest))
	}
	val,_ = url.PathUnescape(val)
	dest[p] = val
}

//如果是连续下标，则视为[]interface{},否则则是map[string]interface{}
func (form *Form) parseArrayItem(dest map[string]interface{}) interface{} {
	len := len(dest)
	var arr []interface{}

	for i := 0; i < len; i++ {
		item, ok := dest[fmt.Sprint(i)]
		if !ok {
			return dest
		}
		arr = append(arr, item)
	}
	return arr
}

func (form *Form) parseArray(dest map[string]interface{}) map[string]interface{} {

	for k, v := range dest {
		mv, ok := v.(map[string]interface{})
		if ok {
			form.parseArray(mv)
			dest[k] = form.parseArrayItem(mv)
		}
	}
	return dest
}
