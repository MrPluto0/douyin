package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsContain(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(IsContain("/login", []string{"/douyin", "/douyin/user", "/douyin/user/register", "/log", "/login"}),
		true, "test IsContain True")

	assert.Equal(IsContain("/login", []string{"/douyin", "/douyin/user", "/douyin/user/register", "/log"}),
		false, "test IsContain False")
}

func BenchmarkIsContain(b *testing.B) {
	arr := []string{"/douyin", "/douyin/user", "/douyin/user/register", "/log", "/login", "/register"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsContain("/login", arr)
	}
}

func Benchmark_IsContain(b *testing.B) {
	arr := []string{"/douyin", "/douyin/user", "/douyin/user/register", "/log", "/login", "/register"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_IsContain("/login", arr)
	}
}
