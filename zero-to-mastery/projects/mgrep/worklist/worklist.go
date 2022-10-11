package worklist

type Entry struct {
	Path string
}

type WorkList struct {
	jobs chan Entry
}

func (w *WorkList) Add(work Entry) {
	w.jobs <- work
}

func (w *WorkList) Next() Entry {
	return <-w.jobs
}

func New(buffSize int) WorkList {
	return WorkList{make(chan Entry, buffSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

func (w *WorkList) Finalise(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
