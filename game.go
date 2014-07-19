package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Game struct {
}

var game = Game{}

type World struct {
	Uptime   int       `json:"uptime"`
	Width    float64   `json:"width"`
	Height   float64   `json:"height"`
	Entities []*Entity `json:"entities"`
}

var (
	world            = World{0, 500.0, 500.0, nil}
	numberOfEntities = 20
	numTeams         = 7
)

func updateGame(now time.Time, w World) World {
	for _, entity := range w.Entities {
		v1 := centerOfMass(entity)
		v2 := separation(entity)
		v3 := neighborVelocity(entity)
		v4 := tendToPlace(entity)

		velocities := []Vector{v1, v2, v3, v4}
		entity.Velocity = entity.Velocity.addMany(velocities)
		entity.Velocity = maxVelocity(entity)
		entity.Position = entity.Position.Add(entity.Velocity)
	}

	out, _ := json.Marshal(w)
	hub.broadcast <- []byte(out)

	return w
}

func (game *Game) run() {
	for i := 0; i < numberOfEntities; i++ {
		id := rand.Intn(numTeams)
		x := world.Width/2 + ((rand.Float64() - 0.5) * 8)
		y := world.Height/2 + ((rand.Float64() - 0.5) * 8)
		ent := Entity{id, Vector{x, y}, Vector{0, 0}}
		world.Entities = append(world.Entities, &ent)
	}
	start := time.Now().Second()
	timer := time.Tick(15 * time.Millisecond)
	for now := range timer {
		world = updateGame(now, world)
		world.Uptime = time.Now().Second() - start
	}
}
