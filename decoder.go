package form

import (
	"errors"
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
	return form.dest, nil
}

func (form *Form) parsePath(path string) ([]string, error) {

	var paths []string
	var current string

	for _, c := range path {
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
