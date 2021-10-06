package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}
type Philo struct {
	id, count           int
	leftChop, rightChop *ChopS
}

func (p Philo) eat(chanPhilos chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		chanPhilos <- &p
		p.leftChop.Lock()
		p.rightChop.Lock()
		fmt.Println("starting to eat", p.id)
		p.count = p.count + 1
		fmt.Println("finishing eating", p.id)
		p.rightChop.Unlock()
		p.leftChop.Unlock()

		wg.Done()

	}
}

// 'host' allows only 2 philosophers to eat at the same time
func host(chanPhilos chan *Philo) {
	for {
		if len(chanPhilos) == 2 {
			<-chanPhilos
			<-chanPhilos

			//time delay for visualising the process of eating
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	var i int
	var wg sync.WaitGroup

	chanPhilos := make(chan *Philo, 2)
	ChopSticks := make([]*ChopS, 5)
	for i = 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([]*Philo, 5)
	for i = 0; i < 5; i++ {
		Philosophers[i] = &Philo{i + 1, 0, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	wg.Add(15)
	go host(chanPhilos)
	for i = 0; i < 5; i++ {
		go Philosophers[i].eat(chanPhilos, &wg)
	}
	wg.Wait()
}
