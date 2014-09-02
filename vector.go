package main

import "math"

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector) addMany(vecs []Vector) Vector {
	for _, vec := range vecs {
		v = v.Add(vec)
	}
	return v
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector) MultiplyByNum(n float64) Vector {
	x := v.X * n
	y := v.Y * n
	return Vector{x, y}
}

func (v Vector) DivideByNum(n float64) Vector {
	return Vector{v.X / n, v.Y / n}
}

func (v Vector) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}
