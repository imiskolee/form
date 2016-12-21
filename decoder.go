package form

import (
	"errors"
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

func (form *Form) Decode() (interface{}, error) {
	form.reset()
	vals, err := url.ParseQuery(form.raw)

	if err != nil {
		return nil, err
	}
	for k, v := range vals {
		paths, err := form.parsePath(k)
		if err != nil {
			return nil, err
		}

		if len(paths) < 1 {
			return nil, errors.New("empty key")
		}

		var c map[string]interface{} = form.dest
		for i := 0; i < len(paths)-1; i++ {
			nv, ok := (c)[paths[i]]
			if ok {
				nm, ok := nv.(map[string]interface{})
				if ok {
					c = nm
				} else {
					return nil, errors.New("must be map[string]interface{}:" + paths[i])
				}
			} else {
				(c)[paths[i]] = make(map[string]interface{})
				nm := (c)[paths[i]].(map[string]interface{})
				c = nm
			}
		}
		if len(v) > 1 {
			(c)[paths[len(paths)-1]] = v
		} else {
			(c)[paths[len(paths)-1]] = v[0]
		}
	}

	dest := form.parseArray(form.dest)
	return dest, nil
}

func (form *Form) parsePath(path string) ([]string, error) {

	var paths []string
	var current string

	for _, c := range path {
		if c == '.' {
			c = '_'
		}
		switch c {
		case '[':
			if len(current) > 0 {
				paths = append(paths, current)
				current = ""
			}
			break
		case ']':
			if len(current) > 0 {
				paths = append(paths, current)
				current = ""
			} else {
				return nil, errors.New("not allow empty key")
			}
		default:
			current += string(c)
		}
	}

	if len(current) > 0 {
		paths = append(paths, current)
		current = ""
	}

	return paths, nil
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

func (form *Form) parseArray(dest map[string]interface{}) interface{} {

	for k, v := range dest {
		mv, ok := v.(map[string]interface{})
		if ok {
			form.parseArray(mv)
			dest[k] = form.parseArrayItem(mv)
		}
	}
	return dest
}
