package main

const (
	StaticEntity = iota
	DynamicEntity
)

// Entity ...
type Entity struct {
	TeamId   int    `json:"team_id"`
	Position Vector `json:"position"`
	Size     Vector `json:"size"`
	Velocity Vector `json:"velocity"`
	Target   Vector `json:"target"`
	Kind     int    `json:"kind"`
}

var (
	separationDistance = 30.0
	jumpSpeed          = 15.0
	velocityLimit      = 5.0
	runSpeed           = 5.0
)

func (e *Entity) Update() {
	if e.Kind == DynamicEntity {
		e.Velocity = e.Velocity.Add(world.Gravity)
		if e.Velocity.X > 0 {
			e.Velocity = e.Velocity.Add(world.Friction)
		} else {
			e.Velocity = e.Velocity.Sub(world.Friction)
		}
		e.ClampMaxVelocity()
		e.Position = e.Position.Sub(e.Velocity)

		// check collisions
		for _, entity := range world.Entities {
			if entity != e {
				e.CheckCollision(entity)
			}
		}
	}
}

func (e *Entity) Jump() {
	e.Velocity = e.Velocity.Add(Vector{0, jumpSpeed})
}

func (e *Entity) Move(dir float64) {
	e.Velocity.X -= runSpeed * dir
}

func (e *Entity) CheckCollision(entity *Entity) {
	// vertical
	if e.Position.Y+e.Size.Y > entity.Position.Y &&
		e.Position.X > entity.Position.X &&
		e.Position.X < entity.Position.X+entity.Size.X {

		e.Velocity.Y = 0
		e.Position.Y = entity.Position.Y - entity.Size.Y
	}
	// horizontal
	//if e.Position.X+e.Size.X > entity.Position.X &&
	//e.Position.Y > entity.Position.Y &&
	//e.Position.Y < entity.Position.Y+entity.Size.Y {

	//e.Velocity.X = 0
	//e.Position.X = entity.Position.X - entity.Size.X
	//}
}

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
			mag := difference.Mag()
			if mag < separationDistance {
				memo = memo.Sub(difference)
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
	return place.Sub(e.Position).divideByNum(100)
}
