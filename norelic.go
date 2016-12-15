package norelic

import (
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

type NoRelicApplication struct{}

type NoRelicTransaction struct {
	http.ResponseWriter
	R *http.Request
}

func (nra NoRelicApplication) StartTransaction(name string, w http.ResponseWriter, r *http.Request) newrelic.Transaction {
	return &NoRelicTransaction{
		w,
		r}
}

func (nra NoRelicApplication) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	return nil
}

func (nt NoRelicTransaction) End() error {
	return nil
}

func (nt NoRelicTransaction) Ignore() error {
	return nil
}

func (nt NoRelicTransaction) SetName(name string) error {
	return nil
}

func (nt NoRelicTransaction) NoticeError(err error) error {
	return nil
}

func (nt NoRelicTransaction) AddAttribute(key string, value interface{}) error {
	return nil
}

func (nt NoRelicTransaction) StartSegmentNow() newrelic.SegmentStartTime {
	return newrelic.SegmentStartTime{}
}
