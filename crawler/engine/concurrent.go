package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// ParseResult channel
	out := make(chan ParseResult)
	e.Scheduler.Run()
	// create goroutines
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	// submit seeds to scheduler
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, reqResult := range result.Requests {
			e.Scheduler.Submit(reqResult)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
