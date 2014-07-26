package main

import "encoding/json"

type Event struct {
	Kind        string
	CommandType string
	Direction   float64
}

func GetInput() {
	for {
		select {
		case m := <-hub.receive:
			rcvdEvents := []Event{}
			err := json.Unmarshal(m, &rcvdEvents)
			if err != nil {
				panic(err)
			}
			for _, e := range rcvdEvents {
				events = append(events, e)
			}
		}
	}
}
