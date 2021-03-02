package kgo

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert_Struct2Map(t *testing.T) {
	//结构体
	var p1 sPerson
	gofakeit.Struct(&p1)
	mp1, _ := KConv.Struct2Map(p1, "json")
	mp2, _ := KConv.Struct2Map(p1, "")

	var ok bool

	_, ok = mp1["name"]
	assert.True(t, ok)

	_, ok = mp1["none"]
	assert.False(t, ok)

	_, ok = mp2["Age"]
	assert.True(t, ok)

	_, ok = mp2["none"]
	assert.True(t, ok)
}

func BenchmarkConvert_Struct2Map_UseTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "json")
	}
}

func BenchmarkConvert_Struct2Map_NoTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "")
	}
}

func TestConver_Int2Str(t *testing.T) {
	var res string

	res = KConv.Int2Str(0)
	assert.NotEmpty(t, res)

	res = KConv.Int2Str(31.4)
	assert.Empty(t, res)

	res = KConv.Int2Str(PKCS_SEVEN)
	assert.Equal(t, "7", res)
}

func BenchmarkConver_Int2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Int2Str(123456789)
	}
}

func TestConver_Float2Str(t *testing.T) {
	var res string

	//小数位为负数
	res = KConv.Float2Str(flPi1, -2)
	assert.Equal(t, 4, len(res))

	res = KConv.Float2Str(flPi2, 3)
	assert.Equal(t, 5, len(res))

	res = KConv.Float2Str(flPi3, 3)
	assert.Equal(t, 5, len(res))

	res = KConv.Float2Str(flPi4, 9)
	assert.Equal(t, 11, len(res))
}

func BenchmarkConver_Float2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Float2Str(flPi2, 3)
	}
}
