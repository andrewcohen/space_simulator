package main

import (
	"encoding/json"
	"log"
	"time"
)

type Game struct {
	World World
}

var game = Game{}

var (
	events           = []Event{}
	numberOfEntities = 10
	numTeams         = 7
)

func (g *Game) Update() {
	timer := time.Tick(15 * time.Millisecond)
	for _ = range timer {
		for _, entity := range g.World.Entities {
			entity.Update(&g.World)
		}
		g.ProcessEvents()

		out, _ := json.Marshal(g.World)
		hub.broadcast <- []byte(out)
	}
}

func (g *Game) AddPlayer(e Event) {
	player := Entity{
		TeamId:   1,
		Mass:     10,
		Size:     Vector{50, 50},
		Position: Vector{200, 10},
	}
	g.World.Entities = append(g.World.Entities, &player)
}

func (g *Game) ProcessEvents() {
	for _, event := range events {
		switch event.CommandType {
		case "join":
			g.AddPlayer(event)
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
				panic(err)
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
	planet := Entity{
		Position: Vector{500, 500},
		Size:     Vector{50, 50},
		Mass:     1e8,
	}
	g.World.Entities = append(g.World.Entities, &planet)

	go g.GetInput()
	go g.Update()

	go func() {
		timer := time.Tick(3 * time.Second)
		for _ = range timer {
			log.Println("Entities: ", len(g.World.Entities))
		}
	}()
}
