package main

import (
	"log"
	"sync"
)

// Engine ...
type Engine struct {
	Vertex    sync.Map
	Pending   chan State
	WaitGroup sync.WaitGroup
}

// Init ...
func (engine *Engine) Init(s State, n int) {
	engine.Pending = make(chan State, n)
	engine.WaitGroup.Add(1)
	engine.Pending <- s
	engine.Vertex.Store(s, MAXSTEP)
}

// Worker ...
func (engine *Engine) Worker(s State) {
	defer engine.WaitGroup.Done()
	for _, t := range s.Moves() {
		val, ok := engine.Vertex.Load(s)
		if ok == false {
			log.Fatalln()
		}
		step := val.(int) + 1
		if t.Final() {
			step = 0
		}
		val, ok = engine.Vertex.Load(t)
		if ok && val.(int) <= step {
			continue
		}
		engine.Vertex.Store(t, step)
		engine.WaitGroup.Add(1)
		engine.Pending <- t
	}
}

// Thread ...
func (engine *Engine) Thread() {
	for s := range engine.Pending {
		engine.Worker(s)
	}
}

// Wait ...
func (engine *Engine) Wait() {
	engine.WaitGroup.Wait()
	close(engine.Pending)
}

// Solve ...
func (engine *Engine) Solve(state State) State {
	now, ok := engine.Vertex.Load(state)
	if ok == false {
		log.Fatalln("[SOLVE]: [INVALID STATE]")
	}

	log.Println("[SOLVE]: [NEED STEP]", now)
	state.Print()

	for _, t := range state.Moves() {
		if val, ok := engine.Vertex.Load(t); ok {
			if val.(int)+1 == now.(int) {
				return t
			}
		}
	}
	log.Fatalln("[SOLVE]: [NO SOLUTION]")
	return State{}
}
