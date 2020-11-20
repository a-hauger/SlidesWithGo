package test

import (
	"errors"
	"testing"
)

type Dog struct {
	lang string
}

func NewDog(l string) *Dog {
	return &Dog{lang: l}
}

func (d *Dog) Bark() (string, error) {
	if d.lang == "english" {
		return "Bark!", nil
	} else if d.lang == "espanol" {
		return "Borf!", nil
	} else {
		return "", errors.New("Language unknown")
	}
}

//START OMIT
func TestNewDog(t *testing.T) {
	arg := "english"
	want := Dog{lang: "english"}
	got := NewDog(arg)

	if *got != want {
		t.Errorf("NewDog(%v) want:%v; got:%v", arg, want, got)
	}
}
func TestDog_Bark(t *testing.T) {
	testDog := NewDog("english")
	want := "Bark!"
	got, err := testDog.Bark()

	if err != nil {
		t.Errorf(err.Error())
	}

	if got != want {
		t.Errorf("Bark() want:%s; got:%s", want, got)
	}
}

//END OMIT
func TestDog_Bark_spanish(t *testing.T) {
	testDog := NewDog("espanol")
	want := "Borf!"
	got, err := testDog.Bark()

	if err != nil {
		t.Errorf(err.Error())
	}

	if got != want {
		t.Errorf("Bark() want:%s; got:%s", want, got)
	}
}
