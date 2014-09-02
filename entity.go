package main

// Entity ...
type Entity struct {
	Position *Vector `json:"position"`
	Velocity *Vector `json:"velocity"`
	Mass     float64 `json:"mass"`
}

const G = 6.67384e-11

func (e *Entity) Update(world *World) {
	// planetary gravity
	force := &Vector{}
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
