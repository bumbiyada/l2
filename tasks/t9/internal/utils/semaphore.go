package utils

type Semaphore interface {
	Down()
	Up()
}

type semaphore struct {
	sem chan struct{}
}

func (s *semaphore) Down() {
	s.sem <- struct{}{}
}

func (s *semaphore) Up() {
	_ = <-s.sem
}

func NewSemaphore(capacity int) Semaphore {
	return &semaphore{
		sem: make(chan struct{}, capacity),
	}
}
