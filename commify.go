package stringutil

import (
	"strconv"
	"strings"
)

func CommifyFloat(f float64, precision int) string {
	var negative = f < 0
	if negative {
		f = -f
	}

	var s = strconv.FormatFloat(f, 'f', precision, 64)

	var end = strings.IndexByte(s, '.')
	if end < 0 {
		end = len(s)
	}

	var rs = []rune(s)
	//rs = rs[:end]

	for end-3 > 0 {
		end -= 3
		//rs = append(rs[:end], append([]rune{','}, rs[end:]...)...)
		rs = append(rs, 0)
		copy(rs[end+1:], rs[end:])
		rs[end] = ','
	}

	if negative {
		rs = append(rs, 0)
		copy(rs[1:], rs)
		rs[0] = '-'
	}

	return string(rs)
}

func CommifyFloat2(f float64, precision int) string {
	var negative = f < 0
	if negative {
		f = -f
	}

	var s = strconv.FormatFloat(f, 'f', precision, 64)

	var end = strings.IndexByte(s, '.')
	if end < 0 {
		end = len(s)
	}

	var rs = []rune(s)
	//rs = rs[:end]

	for end-3 > 0 {
		end -= 3
		rs = append(rs[:end], append([]rune{','}, rs[end:]...)...)
		//rs = append(rs, 0)
		//copy(rs[end+1:], rs[end:])
		rs[end] = ','
	}

	if negative {
		rs = append(rs, 0)
		copy(rs[1:], rs)
		rs[0] = '-'
	}

	return string(rs)
}

func CommifyInt(i int64) string {
	var negative = i < 0
	if negative {
		i = -i
	}

	var s = strconv.FormatInt(i, 10)

	var end = len(s)

	var rs = []rune(s)

	for end-3 > 0 {
		end -= 3
		//rs = append(rs[:end], append([]rune{','}, rs[end:]...)...)
		rs = append(rs, 0)
		copy(rs[end+1:], rs[end:])
		rs[end] = ','
	}

	if negative {
		rs = append(rs, 0)
		copy(rs[1:], rs)
		rs[0] = '-'
	}

	return string(rs)
}
