package popquiz

import "testing"

func Question_One(base int, spend int) int {
	return base - spend
}

func TestQuestion_One(t *testing.T) {
	arg := []int{100, 30}
	want := 70
	got := Question_One(arg[0], arg[1])
	if got != want {
		t.Errorf("Question_One(base=%v, spend=%v) want=%v; got=%v", arg[0], arg[1], want, got)
	}
}
