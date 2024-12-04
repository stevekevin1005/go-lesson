package scheduler

import "go-examples/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) ConfitureMasterWorkerChan(c chan engine.Request) {
	panic("implement me")
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-s.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-s.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
