package main

type World struct {
	Uptime   int       `json:"uptime"`
	Width    float64   `json:"width"`
	Height   float64   `json:"height"`
	Entities []*Entity `json:"entities"`
}

var world = World{0, 500.0, 500.0, nil}
