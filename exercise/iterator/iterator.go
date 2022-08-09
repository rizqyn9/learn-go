package iterator

import (
	"fmt"
)

func For_Loop(callback func(string)) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if i == 5 {
				break
			}
			callback(fmt.Sprint(i, j))
		}
	}
}

func Array() {
	limitedArray := [...]int{1, 2, 3}
	makeArray := make([]string, 5)

	fmt.Println(limitedArray)
	for i, v := range limitedArray {
		fmt.Println(v, i)
	}

	defer func() {
		fmt.Println("Error Limit")
	}()

	makeArray[3] = "asd"

	fmt.Println(len(makeArray))
	fmt.Println(makeArray)
}

func Slice() {
	slice := []string{"asdasd", "slice"}

	fmt.Println(slice)

	for _, data := range slice {
		fmt.Println(data)
	}

	slice2 := append(slice, "Hahah")

	slice3 := slice[1:2]

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func MapExercise() {
	mapData := map[string]int{"data1": 1, "data2": 2}

	fmt.Println(mapData)

	// Check exist
	_, isExist := mapData["data3"]

	fmt.Println(isExist)

	nestedMapData := []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
	}

	fmt.Println(nestedMapData)
}
