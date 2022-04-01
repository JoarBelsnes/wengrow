package bubble


import (
"math/rand"
"testing"
"time"
)

func mockArrayOfInts (size int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(size)
}

func benchmarkBubble(i int, b *testing.B) {
	mockData := mockArrayOfInts(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		BubbleSort(mockData)
	}
}
func BenchmarkBubble10(b *testing.B) { benchmarkBubble(10, b) }
func BenchmarkBubble100(b *testing.B) { benchmarkBubble(100, b) }
func BenchmarkBubble1000(b *testing.B) { benchmarkBubble(1000, b) }
func BenchmarkBubble10000(b *testing.B) { benchmarkBubble(10000, b) }
func BenchmarkBubble100000(b *testing.B) { benchmarkBubble(100000, b) }



