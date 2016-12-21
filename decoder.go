package form

import (
	"fmt"
	"net/url"
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
}

func NewForm(raw string) *Form {
	form := new(Form)
	form.raw = raw
	return form
}

func (form *Form) reset() {
	form.dest = make(map[string]interface{})
}

func (form *Form) Decode() (map[string]interface{}, error) {

	form.reset()

	u, err := url.QueryUnescape(form.raw)

	if err != nil {
		return nil, err
	}

	u = form.raw

	var paths []string

	vals := make(map[string]interface{})

	state := 0

	current := ""

	for _, c := range u {
		switch c {
		case '.':
			//php replace . to _
			if state == 0 {
				c = '_'
			}
		case '[':
			if state == 0 && len(current) > 0 {
				paths = append(paths, current)
				current = ""
				continue
			}
			if state == 0 {
				continue
			}
		case ']':
			if state == 0 {
				paths = append(paths, current)
				current = ""
				continue
			}
		case '&':
			if state == 1 && len(current) > 0 {
				insertValue(&vals, paths, current)
				current = ""
				paths = make([]string, 0)
			}
			state = 0
			continue
		case '=':
			if state == 0 {
				if len(current) > 0 {
					paths = append(paths, current)
					current = ""
				}
				state = 1
				continue
			}
		}
		current = current + string(c)
	}
	if state == 1 && len(current) > 0 {
		insertValue(&vals, paths, current)
		current = ""
		paths = make([]string, 0)
	}
	return form.parseArray(vals), nil
}

func insertValue(destP *map[string]interface{}, path []string, val string) {
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
