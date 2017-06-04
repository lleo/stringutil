/*
Package stringutil implements a few basic functions I've wanted, but are
not great (just good enough for me). `LessThan(string, string) bool` is
particularly at odds with the usual way.


String Incrementor
------------------
This is for producing monotonically increasing strings with only some
subset of alphanumeric letters.

One of the functions is basically ok is `Reverse(string) string`; which
does what you think.

The newest variation on my string incrementor algorithm is controled by a
configuration structure. That config structure is used directly to call
the `cfg.Inc(string) string` method. You can use the method via a cute
Golang trick that takes an object.method call and wraps it in a function:

    var Inc = stringutil.AlphaNum.Inc
    str1 = Inc(str0)

The classic one is `Inc(string) string`, but only works on lower-case
ascii alphabetic letters.

The last incrementor implements it's own type `Str`. It is even more lame
than the others; don't bother. It's only there cuz some old programs of
 mine might still use it.
*/
package stringutil

import (
	"regexp"
)

// StringIncr struct defines the configuration for calls to
// `cfg.Inc(string) string` .
type StringIncr struct {
	first      rune
	last       rune
	notAllowed *regexp.Regexp
}

// AlphaNum is the configuration struct For `Inc(string)` on alphanumeric
// strings.
var AlphaNum = StringIncr{
	first:      '0',
	last:       'z',
	notAllowed: regexp.MustCompile("[^[:alnum:]]+"),
}

// Digit is the configuration struct For `Inc(string)` on all digit strings.
var Digit = StringIncr{
	first:      '0',
	last:       '9',
	notAllowed: regexp.MustCompile("[^[:digit:]]+"),
}

// Alpha is the configuration struct For `Inc(string)` on upper and lower case
// alphabetic only strings.
var Alpha = StringIncr{
	first:      'A',
	last:       'z',
	notAllowed: regexp.MustCompile("[^[:alpha:]]+"),
}

// Upper is the configuration struct For `Inc(string)` on upper-case only
// alphabetic strings.
var Upper = StringIncr{
	first:      'A',
	last:       'Z',
	notAllowed: regexp.MustCompile("[^[:upper:]]+"),
}

// Lower is the configuration struct For `Inc(string)` on lower-case only
// alphabetic strings.
var Lower = StringIncr{
	first:      'a',
	last:       'z',
	notAllowed: regexp.MustCompile("[^[:lower:]]+"),
}

// Inc is the only method for any of the StringInc configuration structs.
// It monotonically increases strings according to my
// `LessThan(string, string) bool`functions constraints.
func (incr StringIncr) Inc(s string) string {
	s = incr.notAllowed.ReplaceAllLiteralString(s, "")
	rs := []rune(s)

	i := 0
	if len(rs) == 0 {
		rs = append(rs, incr.first)
		i++
	}

	l := len(rs)
	p := l - 1
	times := 1
	for i < times {
		switch {
		case rs[p] == '9':
			rs[p] = 'A'
			i++
		case rs[p] == 'Z':
			rs[p] = 'a'
			i++
		case rs[p] == incr.last:
			if p == 0 {
				rs[p] = incr.first
				rs = append([]rune{incr.first}, rs...)
				i++
				l = len(rs)
				p = l - 1
			} else {
				rs[p] = incr.first
				p--
				//i is untouched
			}
		default:
			rs[p]++
			i++
			p = l - 1
		}
	}

	return string(rs)
}

//
// Functional form
//

const _first = 'a' //'0'
const _last = 'z'

// Inc function is the purely functional lower case alpha only that
// obey the constraints of my `LessThan(string, string) bool` function.
// It increments strings like algerbraic numerals (as opposed to to roman
// numerals).
// (Identical to `DigitalInc(string) string`)
func Inc(s string) string {
	// Thae variable "times" used to be an argument to this algorithm, but
	// there still needs to be a loop.
	times := 1

	// Must remove all characters that are not covered by this algorithm.

	rs := []rune(string(s))

	i := 0
	if len(rs) == 0 {
		rs = append(rs, _first)
		i++
	}

	l := len(rs)
	p := l - 1
	for i < times {
		switch {
		case rs[p] == '9':
			rs[p] = 'A'
			p--
		case rs[p] == 'Z':
			rs[p] = 'a'
			p--
		case rs[p] == _last:
			if p == 0 {
				rs[p] = _first
				rs = append([]rune{_first}, rs...)
				i++
				l = len(rs)
				p = l - 1
			} else {
				rs[p] = _first
				p--
				//i is untouched
			}
		default:
			rs[p]++
			i++
			p = l - 1
		}
	}

	return string(rs)
}

// DigitalInc function is the purely functional lower case alpha only that
// obey the constraints of my `LessThan(string, string) bool` function.
// It increments strings like algerbraic numerals (as opposed to to roman
// numerals).
// (Identical to `Inc(string) string`)
func DigitalInc(s string) string {
	// Thae variable "times" used to be an argument to this algorithm, but
	// there still needs to be a loop.
	times := 1

	// Must remove all characters that are not covered by this algorithm.

	rs := []rune(string(s))

	i := 0
	if len(rs) == 0 {
		rs = append(rs, _first)
		i++
	}

	l := len(rs)
	p := l - 1
	for i < times {
		switch {
		case rs[p] == '9':
			rs[p] = 'A'
			p--
		case rs[p] == 'Z':
			rs[p] = 'a'
			p--
		case rs[p] == _last:
			if p == 0 {
				rs[p] = _first
				rs = append([]rune{_first}, rs...)
				i++
				l = len(rs)
				p = l - 1
			} else {
				rs[p] = _first
				p--
				//i is untouched
			}
		default:
			rs[p]++
			i++
			p = l - 1
		}
	}

	return string(rs)
}

// FatInc is the first stupid version of incrementing numbers by my version
// of strings defined with `LessThan(string, string) bool` constraints.
// I only have it cuz it tickles bugs differently than `DigitalInc()`.
// Only works on lower-case ascii letters.
func FatInc(s string) string {
	rs := []rune(string(s))

	i := 0
	if len(rs) == 0 {
		rs = append(rs, _first)
		i++
	}

	p := len(rs) - 1
	times := 1
	for i < times {
		switch {
		case rs[p] == _last:
			p--
		case rs[p] == '9':
			rs[p] = 'A'
			i++
			p = len(rs) - 1
		case rs[p] == 'Z':
			rs[p] = 'a'
			i++
			p = len(rs) - 1
		case p > 0 && rs[p]-rs[p-1] == 1:
			p--
		case p > 0 && rs[p]-rs[p-1] >= 2:
			p--
			fallthrough
		default:
			switch {
			case rs[p] == '9':
				rs[p] = 'A'
			case rs[p] == 'Z':
				rs[p] = 'a'
			default:
				rs[p]++
			}
			i++
			p = len(rs) - 1
		}
		if p == -1 {
			//all runes are _last
			//reset them to _first
			for i := 0; i < len(rs); i++ {
				rs[i] = _first
			}
			//append a new _first; so all
			rs = append(rs, _first)
			//reset the position to the end of the []rune
			p = len(rs) - 1
			i++
		}
	}
	return string(rs)
}

// LessThan function implements a comparison of two strings according to
// rules I prefer to the default meaning implied Golang's by '<' .
//
//Rules for string comparison:
// 0 - shorter is less than longer
// 1 - equal length are compared as strings
func LessThan(s0, s1 string) bool {
	if len(s0) < len(s1) {
		return true
	}
	if len(s0) > len(s1) {
		return false
	}
	//len(s0) == len(s1)
	return string(s0) < string(s1)
}

//
// Immutable Object form
//

// Str type is just an alias for `string` that allows me to hang my own
// string incrementor methods off it.
// It is stupid, just don't use it. It is here for backward compatability
// for any of my old programs that still use it.
type Str string

// Inc method monotonicalll increases the Str returning a new Str. It works
// only on lower-case only ascii alphas.
// It is stupid, just don't use it. It is here for backward compatability
// for any of my old programs that still use it.
func (s Str) Inc(times int) Str {
	rs := []rune(string(s))

	i := 0
	if len(rs) == 0 {
		rs = append(rs, _first)
		i++
	}

	l := len(rs)
	p := l - 1
	for i < times {
		switch {
		case rs[p] == '9':
			rs[p] = 'A'
			p--
		case rs[p] == 'Z':
			rs[p] = 'a'
			p--
		case rs[p] == _last:
			if p == 0 {
				rs[p] = _first
				rs = append([]rune{_first}, rs...)
				i++
				l = len(rs)
				p = l - 1
			} else {
				rs[p] = _first
				p--
				//i is untouched
			}
		default:
			rs[p]++
			i++
			p = l - 1
		}
	}

	return Str(string(rs))
}

// DigitalInc method monotonically increases the Str returning a new Str. It works
// only on lower-case only ascii alphas.
// It is stupid, just don't use it. It is here for backward compatability
// for any of my old programs that still use it.
func (s Str) DigitalInc(times int) Str {
	rs := []rune(string(s))

	i := 0
	if len(rs) == 0 {
		rs = append(rs, _first)
		i++
	}

	l := len(rs)
	p := l - 1
	for i < times {
		switch {
		case rs[p] == '9':
			rs[p] = 'A'
			p--
		case rs[p] == 'Z':
			rs[p] = 'a'
			p--
		case rs[p] == _last:
			if p == 0 {
				rs[p] = _first
				rs = append([]rune{_first}, rs...)
				i++
				l = len(rs)
				p = l - 1
			} else {
				rs[p] = _first
				p--
				//i is untouched
			}
		default:
			rs[p]++
			i++
			p = l - 1
		}
	}

	return Str(string(rs))
}

// FatInc is the first stupid version of incrementing numbers by my version
// of strings defined with `LessThan(string, string) bool` constraints.
// I only have it cuz it tickles bugs differently than `DigitalInc()`.
// Only works on lower-case ascii letters.
func (s Str) FatInc(times int) Str {
	rs := []rune(string(s))

	i := 0
	if len(rs) == 0 {
		rs = append(rs, _first)
		i++
	}

	p := len(rs) - 1
	for i < times {
		switch {
		case rs[p] == _last:
			p--
		case rs[p] == '9':
			rs[p] = 'A'
			i++
			p = len(rs) - 1
		case rs[p] == 'Z':
			rs[p] = 'a'
			i++
			p = len(rs) - 1
		case p > 0 && rs[p]-rs[p-1] == 1:
			p--
		case p > 0 && rs[p]-rs[p-1] >= 2:
			p--
			fallthrough
		default:
			switch {
			case rs[p] == '9':
				rs[p] = 'A'
			case rs[p] == 'Z':
				rs[p] = 'a'
			default:
				rs[p]++
			}
			i++
			p = len(rs) - 1
		}
		if p == -1 {
			//all runes are _last
			//reset them to _first
			for i := 0; i < len(rs); i++ {
				rs[i] = _first
			}
			//append a new _first; so all
			rs = append(rs, _first)
			//reset the position to the end of the []rune
			p = len(rs) - 1
			i++
		}
	}
	return Str(string(rs))
}

// LessThan function implements a comparison of two strings according to
// rules I prefer to the default meaning implied Golang's by '<' .
//
//Rules for Str:
// 0 - shorter is less than longer
// 1 - equal length are compared as strings
func (s Str) LessThan(s1 Str) bool {
	if len(s) < len(s1) {
		return true
	}
	if len(s) > len(s1) {
		return false
	}
	//len(s) == len(s1)
	return string(s) < string(s1)
}

func (s Str) String() string {
	return string(s)
}
