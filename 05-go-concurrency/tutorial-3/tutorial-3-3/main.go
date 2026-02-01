package main

import (
	"sync"
	"time"
)

type People struct {
	name string
}

type Toilet struct {
	available bool
	cond      *sync.Cond
}

func (p *People) UseToilet(t *Toilet, wg *sync.WaitGroup) {
	defer wg.Done()

	t.cond.L.Lock()
	for !t.available {
		t.cond.Wait()
	}

	t.available = false
	time.Sleep(500 * time.Millisecond) // Simulate time taken to use the toilet
	println(p.name, "is using the toilet")

	t.available = true
	t.cond.Signal()
	t.cond.L.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	toilet := Toilet{
		available: true,
		cond:      sync.NewCond(&mu),
	}

	john := People{name: "John"}
	jane := People{name: "Jane"}

	wg.Add(2)
	go john.UseToilet(&toilet, &wg)
	go jane.UseToilet(&toilet, &wg)

	wg.Wait()
}
