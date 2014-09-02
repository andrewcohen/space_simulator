package main

type World struct {
	Width    float32   `json:"width"`
	Height   float32   `json:"height"`
	Gravity  Vector    `json:"gravity"`
	Entities []*Entity `json:"entities"`
}
