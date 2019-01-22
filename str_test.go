package stringutil

import (
	"testing"
)

// s+1 == "0" aka AlphaNum.first
func TestIncOne(t *testing.T) {
	var s string

	s = Inc(s)

	if s != "0" {
		t.Fail()
	}
}

// s+1 == "0" aka AlphaNum.first
// s+62 == "00" aka all of 0..9 then A..Z then a..Z then add a digit
func TestInc63(t *testing.T) {
	var s string

	for i := 0; i < 63; i++ {
		s = Inc(s)
		//fmt.Printf("s++ = %q\n", s)
	}

	if string(s) != "00" {
		t.Fail()
	}

}

// var s = "" aka var s string
// s+1 == "0" aka AlphaNum.first
// s+62 == "00" aka all of 0..9 then A..Z then a..Z then add a digit
// s+62 == "10" aka inc the second digit of the string
func TestInc125(t *testing.T) {
	var s string

	for i := 0; i < 125; i++ {
		s = Inc(s)
		//fmt.Printf("s++ = %q\n", s)
	}

	if string(s) != "10" {
		t.Fail()
	}
}

func TestAlphaNumWrap(t *testing.T) {
	if AlphaNum.Inc("z") != "00" {
		t.Fail()
	}
}

func TestUpperWrap(t *testing.T) {
	if Upper.Inc("Z") != "AA" {
		t.Fail()
	}
}

func TestLowerIncWrap(t *testing.T) {
	if Lower.Inc("z") != "aa" {
		t.Fail()
	}
}

func TestDigitIncWrap(t *testing.T) {
	if Digit.Inc("9") != "00" {
		t.Fail()
	}
}
