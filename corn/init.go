package corn

import (
	"container/heap"
	"time"
)

func NewCorn() *Corn {

	var s Corn
	heap.Init(&s.events)
	return &s

}

func (s *Corn) Push(e Event) {
	heap.Push(&s.events, e)
}

func (s *Corn) Pop() {
	heap.Pop(&s.events)
}

func (s *Corn) Front() Event {
	return s.events[0]
}

func (s *Corn) Reset() {
	s.events = data{}
}

func (s *Corn) AddFunc(t time.Time, f func()) {

	s.Push(Event{
		Time: t,
		Func: f,
	})

}

func (s *Corn) Start() {
	go s.DoWorks()
}

func (s *Corn) DoWorks() {

	for {

		if s.events.Len() == 0 {
			return
		}

		now := s.Front()
		s.Pop()

		time.AfterFunc(time.Duration(now.Time.Unix()-time.Now().Unix()), now.Func)

	}

}
