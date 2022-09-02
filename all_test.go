package calculation

import (
	"fmt"
	"testing"
)

func Test_Abs(t *testing.T) {
	var x, y int
	x = Abs(15)
	y = Abs(-15)

	if x != y {
		t.Error("Didnt expected true")
	}
}

func Test_Min(t *testing.T) {
	var x, y int
	y = 1
	arr := []int{5, 2, 1, 9, 7}
	x = Min(arr...)

	if x != y {
		t.Error("Expected ", y, "got ", x)
	}

	y = 1
	x = Min(1, 9, 8)
	if x != y {
		t.Error("Expected ", y, "got ", x)
	}

	y = 0
	x = Min(1, 3, 0)
	if x != y {
		t.Error("Expected ", y, "got ", x)
	}
}

func Test_Max(t *testing.T) {

	var x, y int
	y = 9
	arr := []int{5, 2, 1, 9, 7}
	x = Max(arr...)

	if x != y {
		t.Error("Expected ", y, "got ", x)
	}

	y = 8
	x = Max(1, 3, 8)
	if x != y {
		t.Error("Expected ", y, "got ", x)
	}

	y = 3
	x = Max(1, 3, 0)
	if x != y {
		t.Error("Expected ", y, "got ", x)
	}
}

func Test_Pow(t *testing.T) {
	var x, y int
	y = 16
	x = Pow(2, 4)
	if x != y {
		t.Error("Expected ", y, "got ", x)
	}
}

func Test_Multiply(t *testing.T) {
	var x, y int
	x = 20
	y = Multiply(4, 5)
	fmt.Println(y)

	if x != y {
		t.Error("Expected ", y, "got ", x)
	}
}
