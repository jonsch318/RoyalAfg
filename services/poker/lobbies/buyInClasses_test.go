package lobbies

import (
	"reflect"
	"testing"
)

func TestGetBuyInClass(t *testing.T) {
	classes := [][2]int{{10, 50}, {51, 150}, {151, 500}}

	tries := [][]int{
		{15, 0},
		{100, 1},
		{480, 2},
	}

	for _, e := range tries {
		re := GetBuyInClass(classes, e[0])
		if re != e[1] {
			t.Errorf("Class did not match expectation, got %d, want: %d.", re, e[1])
		}

	}

}

func TestGetBuyInClassFailing(t *testing.T) {
	classes := [][2]int{{10, 50}, {60, 150}, {151, 500}}

	tries := [][]int{
		{5, -1},
		{501, -1},
		{55, -1},
	}

	for _, e := range tries {
		re := GetBuyInClass(classes, e[0])
		if re != e[1] {
			t.Errorf("Class did not match expectation, got %d, want: %d.", re, e[1])
		}

	}
}

func TestOrder(t *testing.T) {
	classes := [][2]int{{60, 150}, {151, 500}, {10, 50}}
	ex := [][2]int{{10, 50}, {60, 150}, {151, 500}}

	re := OrderBuyInClasses(classes)

	if !reflect.DeepEqual(ex, re) {
		t.Errorf("Ordered not correctly")
	}
}
