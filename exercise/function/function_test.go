package function

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_SameParam(t *testing.T) {
	SameParams(1, 2)
}

func Test_RefParam(t *testing.T) {
	num := 1
	fmt.Println(num)
	RefParam(&num)
	fmt.Println(num)
}

func Test_ReturnPredefined(t *testing.T) {
	data, data2 := ReturnPrefefined()

	fmt.Println(data)
	fmt.Println(data2)
}

func Test_FuncVariadic(t *testing.T) {
	res := FuncVariadic(1, 2, 3, 5, 1, 34, 45, 23)

	fmt.Println(res)
}

func Test_FuncClosure(t *testing.T) {
	res := FuncClosure("Data test")("Test res")

	spew.Dump(res)
}
