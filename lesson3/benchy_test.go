package benchy

import (
	"fmt"
	"strconv"
	"testing"
)

func Benchmark_FmtHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if fmt.Sprintf("%x", 0x12345678) != "12345678" {
			b.Fail()
		}
	}
}

func Benchmark_StrConvHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strconv.FormatUint(0x12345678, 16) != "12345678" {
			b.Fail()
		}
	}
}
