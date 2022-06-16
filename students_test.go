package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var ppl People = People{
	Person{firstName: "Joe",lastName: "Dohn", birthDay:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
	Person{firstName: "Hi",lastName: "Down", birthDay: time.Date(1999, time.November, 10, 23, 0, 0, 0, time.UTC)},
	Person{firstName: "Chriss",lastName: "Masterpiece", birthDay:  time.Date(1998, time.November, 10, 23, 0, 0, 0, time.UTC)},
}

func TestLen(t *testing.T) {
	got := ppl.Len()
	if got != 3 {
			t.Errorf("Len(-1) = %d; want 3", got)
	}
}

func TestLess(t *testing.T) {
	got := ppl.Less(1,2)
	if !got {
			t.Errorf("Len(-1) = %v; want 3", got)
	}
}

func TestSwap(t *testing.T) {
	var pplLocal People = People{
		Person{firstName: "Joe",lastName: "Dohn", birthDay:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		Person{firstName: "Hi",lastName: "Down", birthDay: time.Date(1999, time.November, 10, 23, 0, 0, 0, time.UTC)},
		Person{firstName: "Chriss",lastName: "Masterpiece", birthDay:  time.Date(1998, time.November, 10, 23, 0, 0, 0, time.UTC)},
	}
	pplLocal.Swap(0,1)
	if pplLocal[0].firstName != "Hi" {
			t.Errorf("Wrong info",)
	}
}

////////////////////////////////////////////////////////////////////////


func TestNew(t *testing.T) {
	got, er := New("1 2 3 \n 1 2 3")
	var resData = []int{}

	resData = append(resData, got.data...)
	ans := Matrix{rows: 2, cols: 3, data: []int{1,2,3,1,2,3}}

	if !Equal(resData, ans.data) || er != nil {
			t.Errorf("data is not equal - %v or err has is not nil - %v", !Equal(resData, ans.data), er )
	}
}

func TestSet(t *testing.T) {
	got, er := New("1 2 3 \n 1 2 3")
	var resData = []int{}
	got.Set(0,0,6)

	resData = append(resData, got.data...)
	ans := Matrix{rows: 2, cols: 3, data: []int{6,2,3,1,2,3}}

	if !Equal(resData, ans.data) || er != nil {
			t.Errorf("data is not equal - %v or err has is not nil - %v", !Equal(resData, ans.data), er )
	}
}


func TestRows(t *testing.T) {
	got, er := New("1 2 3 \n 1 2 3")
	var resData = [][]int{{1,2,3}, {1,2,3}}

	if !EqualRows(got.Rows(), resData) || er != nil {
			t.Errorf("data is not equal - %v or err has is not nil - %v", !EqualRows(got.Rows(), resData), er )
	}
}

func TestCols(t *testing.T) {
	got, er := New("1 2 3 \n 1 2 3")
	var resData = [][]int{{1,1},{2,2}, {3,3}}

	if !EqualRows(got.Cols(), resData) || er != nil {
			t.Errorf("data is not equal - %v or err has is not nil - %v", !EqualRows(got.Cols(), resData), er )
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
			return false
	}
	for i, v := range a {
			if v != b[i] {
					return false
			}
	}
	return true
}

func EqualRows(a, b [][]int) bool {
	for i := 0; i < len(a); i++ {
		return Equal(a[i], b[i])
	}
	return true
}
