package iterator_test

import (
	"fmt"
	"testing"

	"learn-go/exercise/iterator"
)

func Test_ForLoop(t *testing.T) {
	iterator.For_Loop(func(s string) {
		fmt.Println(s)
	})
}

func Test_Array(t *testing.T) {
	iterator.Array()
}

func Test_Slice(t *testing.T) {
	iterator.Slice()
}

func Test_MapExercise(t *testing.T) {
	iterator.MapExercise()
}
