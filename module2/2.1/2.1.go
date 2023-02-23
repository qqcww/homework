package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

type goods struct {
	num  int
	cond *sync.Cond
}

var cpuprofile = flag.String("cpuprofile", "./tmp/cpuprofile", "write cpu profile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	wg := sync.WaitGroup{}
	wg.Add(4)
	goods := goods{
		num:  0,
		cond: sync.NewCond(&sync.Mutex{}),
	}
	for i := 1; i < 3; i++ {
		go goods.producer(i, &wg)
	}
	for i := 1; i < 3; i++ {
		go goods.consumer(i, &wg)
	}
	wg.Wait()
}

func (g *goods) producer(n int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		g.cond.L.Lock()
		time.Sleep(3 * time.Second)
		g.num++
		fmt.Printf("%d号厂家生产商品,当前商品数:%d\n", n, g.num)
		g.cond.Signal()
		g.cond.L.Unlock()
	}
	wg.Done()
}

func (g *goods) consumer(n int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		g.cond.L.Lock()
		if g.num <= 0 {
			// fmt.Printf("%d号消费者,货架商品数为%d,等待补货\n", n, g.num)
			g.cond.Wait()
		}
		g.num--
		// fmt.Printf("%d号消费者消费商品,当前商品数:%d\n", n, g.num)
		g.cond.L.Unlock()
	}
	wg.Done()
}
