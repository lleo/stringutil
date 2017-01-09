package stringutil

import (
	"fmt"
	"testing"
)

func TestIncOne(t *testing.T) {
	var s = Str("")

	fmt.Printf("BEFORE: s = %q\n", s)

	s = s.Inc(1)

	fmt.Printf("AFTER: s = %q\n", s)

	if string(s) != "a" {
		t.Fail()
	}
}

func TestInc27(t *testing.T) {
	var s = Str("")

	for i := 0; i < 27; i++ {
		s = s.Inc(1)
		fmt.Printf("s++ = %q\n", s)
	}

	if string(s) != "aa" {
		t.Fail()
	}

}

func TestInc27Alt(t *testing.T) {
	var s = Str("")

	fmt.Printf("BEFORE s=%q\n", s)
	s = s.Inc(27)
	fmt.Printf("AFTER: s = s.Inc(27) = %q\n", s)

	if string(s) != "aa" {
		t.Fail()
	}

}

func TestInc53(t *testing.T) {
	var s = Str("")

	for i := 0; i < 53; i++ {
		s = s.Inc(1)
		fmt.Printf("s++ = %q\n", s)
	}

	if string(s) != "ba" {
		t.Fail()
	}
}

var MAXLEN = 7

func TestIncLen(t *testing.T) {
	var s = Str("")

	//s = s.Inc(1) //get off ""
	s = s.FatInc(1) //get off ""

	var l int
	var lastLen = len(s)
	for i := 0; l < MAXLEN; i++ {
		//s = s.Inc(1)
		s = s.FatInc(1)
		l = len(s)
		if l > lastLen {
			fmt.Printf("s = %s\n", s)
			fmt.Printf("current length = %d; last legth = %d; i = %d\n", l, lastLen, i)
			lastLen = l
		}
	}
}
