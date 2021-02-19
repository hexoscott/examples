package stringperf

import (
	"fmt"
	"strconv"
	"strings"
)

var lookup = map[int]string{}

func init() {
	for i := 0; i < 1000; i++ {
		lookup[i] = strconv.FormatInt(int64(i), 10)
	}
}

func useSprintf() string {
	b := strings.Builder{}

	for i := 0; i < 1000; i++ {
		b.Write([]byte(fmt.Sprintf("%v", i)))
	}

	return b.String()
}

func useItoA() string {
	b := strings.Builder{}

	for i := 0; i < 1000; i++ {
		b.Write([]byte(strconv.Itoa(i)))
	}

	return b.String()
}

func useStrconv() string {
	b := strings.Builder{}

	for i := 0; i < 1000; i++ {
		b.Write([]byte(strconv.FormatInt(int64(i), 10)))
	}

	return b.String()
}

func mapLookup() string {
	b := strings.Builder{}

	for i := 0; i < 1000; i++ {
		b.Write([]byte(lookup[i]))
	}

	return b.String()
}
