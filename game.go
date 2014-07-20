package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Game struct {
}

var game = Game{}

var (
	events           = []Event{}
	numberOfEntities = 10
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

	return w
}

func update() {
	start := time.Now().Second()
	timer := time.Tick(15 * time.Millisecond)
	for now := range timer {
		world = updateGame(now, world)
		processEvents()
		world.Uptime = time.Now().Second() - start

		out, _ := json.Marshal(world)
		hub.broadcast <- []byte(out)
	}
}

func updateTarget(t Vector) {
	for _, entity := range world.Entities {
		entity.Target = t
	}
}

func processEvents() {
	for _, event := range events {
		switch event.Kind {
		case "click":
			updateTarget(Vector{event.X, event.Y})
			break
		}
	}

	events = nil
}

func (game *Game) run() {
	println("run")
	for i := 0; i < numberOfEntities; i++ {
		id := rand.Intn(numTeams)
		x := world.Width/2 + ((rand.Float64() - 0.5) * 8)
		y := world.Height/2 + ((rand.Float64() - 0.5) * 8)
		ent := Entity{id, Vector{x, y}, Vector{0, 0}, Vector{100.0, 100.0}}
		world.Entities = append(world.Entities, &ent)
	}
	go getInput()
	go update()
}
