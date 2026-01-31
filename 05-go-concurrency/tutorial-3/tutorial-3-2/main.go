package main

import "sync"

type counter struct {
	mutex sync.Mutex
	val   int
}

func (c *counter) incrementBy1000000(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mutex.Lock()
	// Critical section
	for i := 0; i < 1000000; i++ {
		c.val++
	}
	c.mutex.Unlock()
}

func (c *counter) incrementBy1000000NoMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		c.val++
	}
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	c_mutex := counter{
		val:   0,
		mutex: mutex,
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go c_mutex.incrementBy1000000(&wg)
	}
	wg.Wait()
	println("c_mutex.val:", c_mutex.val)

	c_no_mutex := counter{
		val:   0,
		mutex: mutex,
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go c_no_mutex.incrementBy1000000NoMutex(&wg)
	}
	wg.Wait()
	println("c_no_mutex.val:", c_no_mutex.val)
}
