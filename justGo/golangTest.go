package main

func Sum(in1 int, in2 int) int {
	return in1 + in2
}

func main() {
	check := Sum(1, 1)
	if check != (1 + 1) {
		panic("EVERYBODY PANIC!!!")
	}
}
