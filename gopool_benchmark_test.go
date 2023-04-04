package main

import (
	"sync"
	"testing"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/lesismal/nbio/taskpool"
	"github.com/panjf2000/ants/v2"
)

const (
	PoolSize   = 1000
	BenchTimes = 10000
)

func BenchmarkGo(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(BenchTimes)
		for j := 0; j < BenchTimes; j++ {
			go func() {
				demoFunc()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkGopool(b *testing.B) {
	p := gopool.NewPool("test", PoolSize, &gopool.Config{})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(BenchTimes)
		for j := 0; j < BenchTimes; j++ {
			p.Go(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func BenchmarkAnts(b *testing.B) {
	p, _ := ants.NewPool(PoolSize)
	defer p.Release()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(BenchTimes)
		for j := 0; j < BenchTimes; j++ {
			p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func BenchmarkNbio(b *testing.B) {
	p := taskpool.NewMixedPool(PoolSize, 1, 10000)
	defer p.Stop()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(BenchTimes)
		for j := 0; j < BenchTimes; j++ {
			p.Go(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func demoFunc() int {
	var sum int
	for i := 0; i < 100; i++ {
		sum += i
	}
	return sum
}
