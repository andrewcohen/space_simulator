package main

import "math"

type Entity struct {
	TeamId   int    `json:"team_id"`
	Position Vector `json:"position"`
	Velocity Vector `json:"velocity"`
	Target   Vector `json:"target"`
}

var (
	separationDistance = 30.0
	velocityLimit      = 3.0
)

func centerOfMass(e *Entity) Vector {
	memo := Vector{0, 0}
	for _, entity := range world.Entities {
		if entity != e {
			memo = memo.Add(e.Position)
		}
	}
	memo = memo.divideByNum(float64(len(world.Entities) - 1))
	center := memo.Sub(e.Position).divideByNum(100.0)

	return center
}

func separation(e *Entity) Vector {
	memo := Vector{0, 0}
	for _, entity := range world.Entities {
		if entity != e {
			difference := entity.Position.Sub(e.Position)
			if difference.Mag() < separationDistance {
				memo = memo.Sub(difference)
			}

			if math.Floor(difference.Mag()) == 0.0 {
				memo = memo.Add(Vector{2, 2})
			}
		}
	}
	return memo
}

func neighborVelocity(e *Entity) Vector {
	memo := Vector{0, 0}
	for _, entity := range world.Entities {
		if e != entity {
			memo = memo.Add(e.Velocity)
		}
	}
	memo = memo.divideByNum(float64(len(world.Entities) - 1))
	velocity := memo.Sub(e.Velocity).divideByNum(8)

	return velocity
}

func maxVelocity(e *Entity) Vector {
	velocity := e.Velocity
	mag := velocity.Mag()
	if mag > velocityLimit {
		velocity = velocity.divideByNum(mag).multiplyByNum(velocityLimit)
	}
	return velocity
}

func tendToPlace(e *Entity) Vector {
	place := e.Target
	return place.Sub(e.Position).divideByNum(100)
}
