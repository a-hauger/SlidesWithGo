package popquiz

import "testing"

func Question_Two(base int, spend int) int {
	return base - spend
}

func TestQuestion_Two(t *testing.T) {
	arg := []int{100, 30}
	want := (arg[0] - arg[1])
	got := Question_Two(arg[0], arg[1])
	if got != want {
		t.Errorf("Question_Two(base=%v, spend=%v) want=%v; got=%v", arg[0], arg[1], want, got)
	}
}
