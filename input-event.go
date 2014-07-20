package main

import "encoding/json"

type Event struct {
	Kind string
	X    float64
	Y    float64
}

func getInput() {
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
