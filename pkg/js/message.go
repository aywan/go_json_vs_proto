package js

import "time"

type Attachments struct {
	ID   int32  `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Data []byte `json:"data,omitempty"`
}

type Message struct {
	ID      string        `json:"id,omitempty"`
	Time    time.Time     `json:"time"`
	Message string        `json:"message,omitempty"`
	Status  string        `json:"status,omitempty"`
	Attach  []Attachments `json:"attach,omitempty"`
	Counter int64         `json:"counter,omitempty"`
}
