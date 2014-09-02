package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Game struct {
	World     World
	StepCount int
}

var game = Game{}

var (
	events           = []Event{}
	numberOfEntities = 10
	numTeams         = 7
)

func (g *Game) Update() {
	timer := time.Tick(10 * time.Millisecond)
	for _ = range timer {
		for _, entity := range g.World.Entities {
			entity.Update(&g.World)
		}
		g.ProcessEvents()

		out, _ := json.Marshal(g.World)
		hub.broadcast <- []byte(out)
		g.StepCount++
	}
}

func Rand(max float64) float64 {
	return rand.Float64() * max
}

func RandomPlanet() *Entity {
	return &Entity{
		Mass:     Rand(10000),
		Position: &Vector{Rand(1000), Rand(1000)},
		Velocity: &Vector{0, 0},
	}
}

func (g *Game) AddEntity(e *Entity) {
	g.World.Entities = append(g.World.Entities, e)
}

func (g *Game) ProcessEvents() {
	for _, event := range events {
		switch event.CommandType {
		case "add_planet":
			g.AddEntity(RandomPlanet())
			break
		}
	}

	events = nil
}

type Event struct {
	Kind        string
	CommandType string
	Direction   float64
}

func (g *Game) GetInput() {
	for {
		select {
		case m := <-hub.receive:
			rcvdEvents := []Event{}
			err := json.Unmarshal(m, &rcvdEvents)
			if err != nil {
				log.Fatalln("Error in GetInput:", err)
			}
			for _, e := range rcvdEvents {
				events = append(events, e)
			}
		}
	}
}

func (g *Game) Run() {
	g.World = World{
		Width:   500.0,
		Height:  500.0,
		Gravity: Vector{0, 0},
	}

	go g.GetInput()
	go g.Update()

	go func() {
		timer := time.Tick(3 * time.Second)
		for _ = range timer {
			log.Println("steps/sec:", g.StepCount/3)
			g.StepCount = 0
			log.Println("Entities: ", len(g.World.Entities))
		}
	}()
}
