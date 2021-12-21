package main

import (
	"crypto/md5"
	"testing"
)

func BenchmarkXxh3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xxh3Hash(data)
	}
}

func BenchmarkCityhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cityHash(data)
	}
}

func BenchmarkT1ha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t1haHash(data)
	}
}

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5.Sum(data)
	}
}
