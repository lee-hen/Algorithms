package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	leftCS, rightCS *Chopstick
}

var wg sync.WaitGroup

// All philosophers might lock their left chopsticks concurrently
// All chopsticks would be locked
// None can lock their right chopsticks
func (p Philosopher) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("eating...")
		p.rightCS.Unlock()
		p.leftCS.Unlock()
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
		philosophers[i] = &Philosopher{chopsticks[i], chopsticks[(i+1)%5]}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}
	wg.Wait()
}
