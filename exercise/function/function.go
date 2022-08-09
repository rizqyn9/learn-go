package function

import "fmt"

func SameParams(num1, num2 int) {
	fmt.Println(num1)
	fmt.Println(num2)
}

func RefParam(ref *int) {
	*ref = 10
}

func ReturnPrefefined() (data string, data2 int) {
	data = "asdad"
	data2 = 10

	return
}

func FuncVariadic(data ...int) float64 {
	var total int
	for _, num := range data {
		total += num
	}
	return float64(total) / float64(len(data))
}

func FuncClosure(data string) func(string) string {
	return func(input string) string {
		return input + data
	}
}
