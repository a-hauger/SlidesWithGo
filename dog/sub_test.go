package test

import "testing"

func TestBark(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want string
	}{
		{"English", "english", "Bark!"},
		{"Spanish", "espanol", "What is the meaning of it all?"},
		{"Swedish", "swedish", "Bork!"},
	}
	//START OMIT
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testDog := NewDog(tc.arg)
			got, err := testDog.Bark()
			if err != nil {
				t.Fatalf("I'm just showing off")
			}
			if got != tc.want {
				t.Errorf("This failed")
			}
		})
	}
}

//END OMIT
