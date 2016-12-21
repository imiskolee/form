package form

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_RawQueryCase1(t *testing.T) {

	query := `query_dsl[query][bool][must][0][term][subscribe]=1&query_dsl[query][bool][must][1][terms][tag_list][]=802343393920557056&query_dsl[size]=5&query_dsl[from]=0&query_dsl[sort][subscribe_dateline]=desc`

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
