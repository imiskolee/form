package form

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_RawQueryCase1(t *testing.T) {

	query := `password=wxhousite!%40%23&email=shenlin%40weixinhost.com`

	decoder := NewForm(query)

	dest, err := decoder.Decode()

	dd, _ := json.Marshal(dest)

	fmt.Println(query, string(dd), err)
}

func Test_RawQueryCase2(t *testing.T) {

	query := `a=1&a=2&a=3`

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
