package csv

import (
	"encoding/csv"
	"io"
)

type Writer interface {
	Write([]string) error
}

type WriteFunc func(record []string) error

func (fn WriteFunc) Write(record []string) error {
	return fn(record)
}

type Syncer struct {
	reader io.Reader
	doneC  chan struct{}
	closeC chan struct{}
}

func NewSyncer(r io.Reader) *Syncer {
	return &Syncer{
		reader: r,
		doneC:  make(chan struct{}, 0),
		closeC: make(chan struct{}, 0),
	}
}

func (s *Syncer) SyncTo(w Writer) error {
	defer close(s.doneC)
	r := csv.NewReader(s.reader)
	for i := 0; ; i++ {
		select {
		case <-s.closeC:
			return nil
		default:
			line, err := r.Read()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			if i == 0 { // skip header
				continue
			}
			if err := w.Write(line); err != nil {
				return err
			}
		}
	}
}

// Close stops the syncer
func (s *Syncer) Close() {
	close(s.closeC)
	<-s.doneC
}
