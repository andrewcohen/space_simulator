package main

import (
	"encoding/json"
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
		//velocities := []Vector{
		//centerOfMass(entity),
		//separation(entity),
		//neighborVelocity(entity),
		//tendToPlace(entity),
		//}
		//entity.Velocity = entity.Velocity.addMany(velocities)
		//entity.Velocity = ClampMaxVelocity(entity)
		//entity.Position = entity.Position.Add(entity.Velocity)
		entity.Update()
	}

	return w
}

func update() {
	timer := time.Tick(15 * time.Millisecond)
	for now := range timer {
		world = updateGame(now, world)
		ProcessEvents()

		out, _ := json.Marshal(world)
		hub.broadcast <- []byte(out)
	}
}

func updateTarget(t Vector) {
	for _, entity := range world.Entities {
		entity.Target = t
	}
}

func AddPlayer(e Event) {
	player := Entity{
		TeamId:   1,
		Kind:     DynamicEntity,
		Size:     Vector{50, 50},
		Position: Vector{200, 10},
	}
	world.Entities = append(world.Entities, &player)
}

func CommandEntity(e Event) {
	for _, entity := range world.Entities {
		switch e.Kind {
		case "jump":
			entity.Jump()
			break
		case "move":
			entity.Move(e.Direction)
			break
		}
	}
}

func ProcessEvents() {
	for _, event := range events {
		switch event.CommandType {
		case "direct":
			CommandEntity(event)
			break
		case "join":
			AddPlayer(event)
			break
		}
	}

	events = nil
}

func (game *Game) run() {
	println("run")
	floor := Entity{
		Kind:     StaticEntity,
		Position: Vector{000, 500},
		Size:     Vector{1000, 50},
	}
	lWall := Entity{
		Kind:     StaticEntity,
		Position: Vector{050, 400},
		Size:     Vector{50, 80},
	}
	rWall := Entity{
		Kind:     StaticEntity,
		Position: Vector{600, 400},
		Size:     Vector{50, 80},
	}
	world.Entities = append(world.Entities, &floor)
	world.Entities = append(world.Entities, &rWall)
	world.Entities = append(world.Entities, &lWall)
	go GetInput()
	go update()
}
