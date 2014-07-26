package main

type World struct {
	Width    float32   `json:"width"`
	Height   float32   `json:"height"`
	Gravity  Vector    `json:"gravity"`
	Friction Vector    `json:"friction"`
	Entities []*Entity `json:"entities"`
}

var world = World{
	Width:    500.0,
	Height:   500.0,
	Gravity:  Vector{0, -0.81},
	Friction: Vector{-0.4, 0},
}
