package sleepsort

import (
	"sync"
	"time"
)

type sorter struct {
	unsorted []int
	sorted   chan int
	waiter   *sync.WaitGroup
}

func NewSorter(unsorted ...int) *sorter {
	waiter := &sync.WaitGroup{}
	waiter.Add(len(unsorted))
	return &sorter{
		unsorted: unsorted,
		sorted:   make(chan int),
		waiter:   waiter,
	}
}
func (s *sorter) Sorted() (result []int) {
	for value := range s.Sort() {
		result = append(result, value)
	}
	return result
}
func (s *sorter) Sort() chan int {
	go s.sort()
	go s.wait()
	return s.sorted
}
func (s *sorter) sort() {
	for _, unsorted := range s.unsorted {
		go s.sleep(unsorted)
	}
}
func (s *sorter) sleep(value int) {
	time.Sleep(time.Second * time.Duration(value))
	s.sorted <- value
	s.waiter.Done()
}
func (s *sorter) wait() {
	s.waiter.Wait()
	close(s.sorted)
}
