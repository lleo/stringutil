package stringutil_test

import (
	"math"
	"testing"

	"github.com/lleo/stringutil"
)

func TestCommifyFloat(t *testing.T) {
	var s = stringutil.CommifyFloat(math.Pi*1e6, 4)
	if s != "3,141,592.6536" {
		t.Fatalf("s,%q != \"3,141,592.6536\"", s)
	}
}

func TestCommifyFloat2(t *testing.T) {
	var s = stringutil.CommifyFloat2(math.Pi*1e6, 4)
	if s != "3,141,592.6536" {
		t.Fatalf("s,%q != \"3,141,592.6536\"", s)
	}
}

func TestCommifyInt(t *testing.T) {
	// This fails with "constant 3.14159e+15 truncated to integer"
	//var i = int64(math.Pi * 1e15)

	// This works but annoys me, cuz it uses two variables.
	//var f float64 = math.Pi * 1e15
	//var i = int64(f)

	// This is still annoying me, cuz it uses a variable
	//var i = int64(float64(math.Pi * 1e15))

	//var s = stringutil.CommifyInt(i)

	var s = stringutil.CommifyInt(int64(float64(math.Pi * 1e15)))
	if s != "3,141,592,653,589,793" {
		t.Fatalf("s,%q != \"3,141,592,653,589,793\"", s)
	}
}

func BenchmarkCommifyFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = stringutil.CommifyFloat(math.Pi*1e15, 4)
		if s != "3,141,592,653,589,793.0000" {
			b.Fatalf("s,%q != \"3,141,592,653,589,793.0000\"", s)
		}
	}
}

func BenchmarkCommifyFloat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = stringutil.CommifyFloat2(math.Pi*1e15, 4)
		if s != "3,141,592,653,589,793.0000" {
			b.Fatalf("s,%q != \"3,141,592,653,589,793.0000\"", s)
		}
	}
}
