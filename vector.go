package main

import "math"

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v *Vector) Sub(v2 *Vector) *Vector {
	return &Vector{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v *Vector) MultiplyByNum(n float64) *Vector {
	return &Vector{v.X * n, v.Y * n, v.Z * n}
}

func (v *Vector) DivideByNum(n float64) *Vector {
	return &Vector{v.X / n, v.Y / n, v.Z / n}
}

func (v *Vector) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}
