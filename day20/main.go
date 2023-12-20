package main

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"strings"
)

type ModuleType int

const (
	broadcaster ModuleType = iota
	flipFlop
	conjunction
)

type Module struct {
	name         string
	moduleType   ModuleType
	destinations []string
	state        bool
	history      map[string]PulseType
}

type ModuleState int
type PulseType int

const (
	Low PulseType = iota
	High
)

func createModule(name string, moduleType ModuleType, destinations []string) Module {
	return Module{name: name, moduleType: moduleType, destinations: destinations, state: false, history: make(map[string]PulseType)}
}

func processLine(input string) Module {
	// split on " -> "
	result := strings.Split(input, " -> ")

	// parse the destinations
	destinationString := result[1]
	destinations := strings.Split(destinationString, ", ")
	// trim the spaces
	for i, destination := range destinations {
		destinations[i] = strings.TrimSpace(destination)
	}

	name := result[0]
	switch name[0] {
	case '%':
		return createModule(name[1:], flipFlop, destinations)
	case '&':
		return createModule(name[1:], conjunction, destinations)
	default:
		return createModule(name, broadcaster, destinations)
	}
}

type Pulse struct {
	source      string
	destination string
	value       PulseType
}

func (p Pulse) String() string {
	valueString := "Low"
	if p.value == High {
		valueString = "High"
	}
	return fmt.Sprintf("%s from %s to %s", valueString, p.source, p.destination)
}

type PulseQueue struct {
	queue      []Pulse
	highPulses int
	lowPulses  int
}

func (p *PulseQueue) send(source string, destinations []string, value PulseType) {
	for _, destination := range destinations {
		if value == High {
			p.highPulses++
		} else {
			p.lowPulses++
		}
		pulse := Pulse{source, destination, value}
		// fmt.Printf("Sending: %v\n", pulse)
		p.queue = append(p.queue, pulse)
	}
}

func (p *PulseQueue) next() Pulse {
	pulse := p.queue[0]
	p.queue = p.queue[1:]
	return pulse
}

func (p *PulseQueue) len() int {
	return len(p.queue)
}

func Part1(filePath string, buttonPresses int) int {
	lines := aoc.ReadFileToLines(filePath)
	modules := make(map[string]Module)

	for _, line := range lines {
		module := processLine(line)
		modules[module.name] = module
	}

	queue := PulseQueue{queue: make([]Pulse, 0)}
	for i := 0; i < buttonPresses; i++ {
		queue.send("button", []string{"broadcaster"}, Low)
	}

	for queue.len() > 0 {
		pulse := queue.next()
		module := modules[pulse.destination]
		switch module.moduleType {
		case broadcaster:
			// fmt.Printf("Broadcaster: %v -> %v\n", pulse, module)
			queue.send(module.name, module.destinations, pulse.value)
		case flipFlop:
			if pulse.value == Low {
				if !module.state {
					queue.send(module.name, module.destinations, High)
				} else {
					queue.send(module.name, module.destinations, Low)
				}
				// copy the module and reverse the state
				modules[pulse.destination] = Module{
					name:         module.name,
					moduleType:   module.moduleType,
					destinations: module.destinations,
					state:        !module.state,
					history:      module.history,
				}
			}
		case conjunction:
			// fmt.Printf("Conjunction: %v -> %v\n", pulse, module)
			modules[pulse.destination].history[pulse.source] = pulse.value
			pulseValue := Low
			for _, next := range module.history {
				if next == Low {
					pulseValue = High
					break
				}
			}
			queue.send(module.name, module.destinations, pulseValue)
		}
	}

	return queue.highPulses * queue.lowPulses
}
