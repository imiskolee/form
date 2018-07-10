package form

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_RawQueryCase1(t *testing.T) {

	query := `query_dsl%5Bquery%5D%5Bbool%5D%5Bmust%5D%5B0%5D%5Bterm%5D%5Bsubscribe%5D=1&query_dsl%5Bquery%5D%5Bbool%5D%5Bmust%5D%5B1%5D%5Bmatch%5D%5Bcomment.mobile%5D=18513502886&query_dsl%5Bsize%5D=50&query_dsl%5Bfrom%5D=0&query_dsl%5Bsort%5D%5Bsubscribe_dateline%5D=desc`

	decoder := NewForm(query)
	dest, err := decoder.Decode()
	dd, _ := json.Marshal(dest)
	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase2(t *testing.T) {

	query := `a=a&a=2&a=3`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase3(t *testing.T) {

	query := `a=1&a[b]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase4(t *testing.T) {

	query := `a[b][c]=1&a[b][c]=1&a[b][d]=1&f=1&g[h][i][l]=1&g[h][i][t]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase5(t *testing.T) {

	query := `a[0]=1&a[1]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase6(t *testing.T) {

	query := `a[b][0]=1&a[b][1]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase7(t *testing.T) {

	query := `a[f]=1&a[b][c][0]=1&a[b][c][1]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase8(t *testing.T) {

	query := `a[]=1&a[]=2&a[]=3`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase10(t *testing.T) {

	query := `a[b][]=1&a[b][]=2&a[b][]=3`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase9(t *testing.T) {

	query := `a[b][]=1&a[c][]=1&e[f][g][]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase11(t *testing.T) {

	query := `a[][a]=1&a[][a]=1&a[][a]=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase12(t *testing.T) {

	query := `a=&b=1`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase13(t *testing.T) {

	query := `token=39b2e29403e8232f6941a2a09c1dfb43&url=http%3A%2F%2Fmp.weixinhost.com%2Faddon%2Fclient_api%3Fa%3Dresource_record%26code%3Da1eaa3ad381aaa768e78d6eb952229d3`
	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase14(t *testing.T) {

	query := `a==b&c=d&e&f====g&g[a]=a`
	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}


func Test_RawQueryCase15(t *testing.T) {

	query := `a[][a]=1&a[][b]=2&a[][a]=3&a[0][c][]=4&a[0][c][]=5`
	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}