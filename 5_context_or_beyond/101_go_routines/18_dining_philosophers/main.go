package main

import (
	"fmt"
	"sync"
)

var host = make(chan bool, 2)
var wg sync.WaitGroup

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	idx int
	leftCS, rightCS *Chopstick
}

// All philosophers might lock their left chopsticks concurrently
// All chopsticks would be locked
// No one can lock their right chopsticks
func (p Philosopher) eat() {
	for {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println(p.idx, "is eating...")
		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
	wg.Done()
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i,chopsticks[i], chopsticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat()
	}
	wg.Wait()
}
