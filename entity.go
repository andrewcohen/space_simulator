package main

import "math"

// Entity ...
type Entity struct {
	TeamId   int     `json:"team_id"`
	Position Vector  `json:"position"`
	Size     Vector  `json:"size"`
	Velocity Vector  `json:"velocity"`
	Target   Vector  `json:"target"`
	Mass     float64 `json:"mass"`
}

const (
	G                  = 6.67384e-11
	separationDistance = 30.0
	velocityLimit      = 5.0
)

func (e *Entity) Update(world *World) {
	// planetary gravity
	force := Vector{}
	for _, entity := range world.Entities {
		if entity != e {
			// F = G*m1*m2*r^2
			d := entity.Position.Sub(e.Position)
			f := d.MultiplyByNum(-G * entity.Mass * e.Mass).DivideByNum(d.Mag())
			force = force.Add(f)
		}
	}
	e.Velocity = e.Velocity.Add(force)
	e.Position = e.Position.Sub(e.Velocity)
}

func (e *Entity) Distance(entity *Entity) Vector {
	x := math.Sqrt(e.Position.X*e.Position.X + entity.Position.X*entity.Position.X)
	y := math.Sqrt(e.Position.Y*e.Position.Y + entity.Position.Y*entity.Position.Y)
	return Vector{x, y}
}

//func centerOfMass(e *Entity) Vector {
//memo := Vector{0, 0}
//for _, entity := range world.Entities {
//if entity != e {
//memo = memo.Add(e.Position)
//}
//}
//memo = memo.DivideByNum(float64(len(world.Entities) - 1))
//center := memo.Sub(e.Position).DivideByNum(100.0)

//return center
//}

//func separation(e *Entity) Vector {
//memo := Vector{0, 0}
//for _, entity := range world.Entities {
//if entity != e {
//difference := entity.Position.Sub(e.Position)
//mag := difference.Mag()
//if mag < separationDistance {
//memo = memo.Sub(difference)
//}
//}
//}

//return memo
//}

//func neighborVelocity(e *Entity) Vector {
//memo := Vector{0, 0}
//for _, entity := range world.Entities {
//if e != entity {
//memo = memo.Add(e.Velocity)
//}
//}
//memo = memo.DivideByNum(float64(len(world.Entities) - 1))
//velocity := memo.Sub(e.Velocity).DivideByNum(8)

//return velocity
//}

func (e *Entity) ClampMaxVelocity() {
	if e.Velocity.X > velocityLimit {
		e.Velocity.X = velocityLimit
	}
	if e.Velocity.Y > velocityLimit {
		e.Velocity.Y = velocityLimit
	}
}

func tendToPlace(e *Entity) Vector {
	place := e.Target
	return place.Sub(e.Position).DivideByNum(100)
}
