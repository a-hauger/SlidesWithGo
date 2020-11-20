package exampleAsATest

import (
	"fmt"
	"testing"
)

func HelloWorld(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func ExampleHelloWorld(t *testing.T) {
	returnedString := HelloWorld("Quick Quack!")
	fmt.Sprintf(returnedString)
	// Output: Hello, Quick Quack!
}
