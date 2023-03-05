package jscw

// Worker ...
type Worker struct{}

// NewWorker ...
func NewWorker() *Worker {
	w := new(Worker)

	return w
}

// Dispose ...
func (w *Worker) Dispose() {}
