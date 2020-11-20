package parallel

import (
	"testing"
	"time"
)

//START OMIT
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}
func TestB(t *testing.T) {
	time.Sleep(5 * time.Second)
}
func TestC(t *testing.T) {
	t.Parallel()
}

//END OMIT
