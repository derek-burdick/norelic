package norelic

import (
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

type application struct{}

type transaction struct {
	http.ResponseWriter
	R *http.Request
}

func NewApplication() newrelic.Application {
	return &application{}
}

func (nra application) StartTransaction(name string, w http.ResponseWriter, r *http.Request) newrelic.Transaction {
	return &transaction{
		w,
		r}
}

func (nra application) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	return nil
}

func (nt transaction) End() error {
	return nil
}

func (nt transaction) Ignore() error {
	return nil
}

func (nt transaction) SetName(name string) error {
	return nil
}

func (nt transaction) NoticeError(err error) error {
	return nil
}

func (nt transaction) AddAttribute(key string, value interface{}) error {
	return nil
}

func (nt transaction) StartSegmentNow() newrelic.SegmentStartTime {
	return newrelic.SegmentStartTime{}
}
