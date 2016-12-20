#### Form 

用于支持PHP的$_POST解析。得到与php版本行为基本一致的map[string]interface{}。

仅用于php系统向GO系统迁移的过渡使用。


#### Examples

```go

       	query := `a[b][c]=1&a[b][c]=1&a[b][d]=1&f=1&g[h][i]=1&g[h][i][t]=1`

       	decoder := form.NewForm(query)

       	dest, err := decoder.Decode()   //dest 是 map[string]interface{}

       	dd, _ := json.Marshal(dest)

       	fmt.Println(query, string(dd), err)
```

  
func NewForm(raw string) *Form

func (form *Form) Decode() (map[string]interface{}, error)
