package stringperf

import "testing"

var result string

func Benchmark_useSprintf(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = useSprintf()
	}

	result = r
}

func Benchmark_useItoA(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = useItoA()
	}

	result = r
}

func Benchmark_useStrconv(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = useStrconv()
	}

	result = r
}

func Benchmark_useMap(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = mapLookup()
	}

	result = r
}
