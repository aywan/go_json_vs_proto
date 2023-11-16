package main

import (
	"encoding/json"
	"testing"
	"time"

	"json_vs_proto/pkg/js"
)

func BenchmarkMarshalJson(b *testing.B) {
	message := js.Message{
		ID:      "d0ff16bc-2d3f-40ce-8807-66d3e331045c",
		Time:    time.Now(),
		Message: "42cf84bc-772e-4094-997c-e0842ac6d835",
		Status:  "c9e0cc36-ed2c-49ed-b0d5-68a6b7b330b4",
		Attach: []js.Attachments{
			{
				ID:   1,
				Type: "0fdb84c1-e121-43a1-917c-892b0badf029",
				Name: "dbaec68d-b5cf-4eb1-9fa2-9ff03803e3e5",
				Data: []byte("9b87f0bb-fd1e-4ed8-84a5-d251e5be1a96"),
			},
			{
				ID:   2,
				Type: "696addee-5b51-4e56-995b-41d5cc8d781d9",
				Name: "3fb1b73d-b218-424c-957b-8b678bf74937",
				Data: []byte("a1cf580c-eab2-40ae-910b-cf3ac801dd2b"),
			},
			{
				ID:   3,
				Type: "bef122d3-4e05-4719-b492-6ceeb9563177",
				Name: "e393e940-9896-4413-bbeb-dd4873c24556",
				Data: []byte("c9ba5fa1-7c82-4a68-8bd7-221ddf849a20"),
			},
		},
		Counter: 0,
	}

	data, err := json.Marshal(message)
	if err != nil {
		b.Fatalf("Failed to marshal json: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err = benchingJson(data)
		if err != nil {
			b.Fatalf("Failed to marshal json message: %v", err)
		}
	}
}

func benchingJson(data []byte) ([]byte, error) {
	var m js.Message

	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	m.Time = time.Now()
	m.Counter = m.Counter + 1

	return json.Marshal(m)
}
