package parallel

import (
	"testing"
	"time"
)

//START OMIT
func TestD(t *testing.T) {
	t.Parallel()
	t.Run("Parent Parallel", func(t *testing.T) {
		time.Sleep(2 * time.Second)
	})
}
func TestE(t *testing.T) {
	time.Sleep(5 * time.Second)
}

func TestF(t *testing.T) {
	t.Parallel()
	defer t.Errorf("Will error before the subtest runs")
	t.Run("Parent Parallel and Subtest Parallel", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
	})
}

//END OMIT
