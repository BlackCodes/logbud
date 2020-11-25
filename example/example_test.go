package example

import "testing"

func TestLogrus(b *testing.T)  {
	Logrus()
}
func BenchmarkLogbud(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sysLog()
	}
}

func BenchmarkLogrus(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i:=0;i<b.N;i++{
		Logrus()
	}
}

func BenchmarkZeroUnCaller(b *testing.B)  {
	b.ResetTimer()
	b.ReportAllocs()
	for i:=0;i<b.N;i++{
		zeroLogs()
	}
}

func BenchmarkZeroCaller(b *testing.B)  {
	b.ResetTimer()
	b.ReportAllocs()
	for i:=0;i<b.N;i++{
		zeroLogsCaller()
	}
}