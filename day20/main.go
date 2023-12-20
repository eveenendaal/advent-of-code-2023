package main

import (
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"strings"
)

type ModuleType int

const (
	broadcaster ModuleType = iota
	flipFlop
	conjunction
	untyped
)

type Module struct {
	name         string
	moduleType   ModuleType
	destinations []string
	state        bool
	history      map[string]PulseType
}

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

type PulseCounter map[PulseType]int64

func (p *PulseCounter) add(pulseType PulseType) {
	(*p)[pulseType]++
}

func (p *PulseCounter) total() int64 {
	return (*p)[Low] * (*p)[High]
}

func SendPulse(pulseCounter PulseCounter, pulses []Pulse, modules map[string]Module) {
	newPulses := []Pulse{}

	for _, pulse := range pulses {
		pulseCounter.add(pulse.value)
		module := modules[pulse.destination]

		switch module.moduleType {
		case broadcaster:
			for _, destination := range module.destinations {
				newPulses = append(newPulses, Pulse{pulse.destination, destination, pulse.value})
			}
		case flipFlop:
			if pulse.value == Low {
				var nextPulse PulseType
				if module.state {
					nextPulse = Low
				} else {
					nextPulse = High
				}
				module.state = !module.state
				modules[pulse.destination] = module

				for _, destination := range module.destinations {
					newPulses = append(newPulses, Pulse{pulse.destination, destination, nextPulse})
				}
			}
		case conjunction:
			module.history[pulse.source] = pulse.value
			modules[pulse.destination] = module

			pulseValue := Low
			for _, next := range module.history {
				if next == Low {
					pulseValue = High
					break
				}
			}

			for _, destination := range module.destinations {
				newPulses = append(newPulses, Pulse{pulse.destination, destination, pulseValue})
			}
		}
	}

	if len(newPulses) > 0 {
		SendPulse(pulseCounter, newPulses, modules)
	}
}

func Part1(filePath string, buttonPresses int) int64 {
	lines := aoc.ReadFileToLines(filePath)
	modules := make(map[string]Module)

	for _, line := range lines {
		module := processLine(line)
		modules[module.name] = module
	}

	for _, module := range modules {
		for _, destination := range module.destinations {
			existingModule, ok := modules[destination]
			if ok {
				existingModule.history[module.name] = Low
				modules[destination] = existingModule
			} else {
				newModule := createModule(destination, untyped, []string{})
				newModule.history[module.name] = Low
				modules[destination] = newModule
			}
		}
	}

	pulseCounter := PulseCounter{}

	for i := 0; i < buttonPresses; i++ {
		SendPulse(pulseCounter, []Pulse{{"button", "broadcaster", Low}}, modules)
	}
	return pulseCounter.total()
}
