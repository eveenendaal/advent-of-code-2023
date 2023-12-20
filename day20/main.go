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

func SendPulse(pulseCounter PulseCounter, pulses []Pulse, modules map[string]Module, metadata *DestinationCycleData) {
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
				if metadata != nil {
					// Is this destination one of the targets we want to be high?
					if _, ok := metadata.targets[pulse.destination]; ok && pulseValue == metadata.expectedPulseType {
						// Record key presses
						metadata.counters[pulse.destination] = metadata.cycles
						// Remove this destination from the targets
						delete(metadata.targets, pulse.destination)
					}

					if len(metadata.targets) == 0 {
						return
					}
				}

				newPulses = append(newPulses, Pulse{pulse.destination, destination, pulseValue})
			}
		}
	}

	if len(newPulses) > 0 {
		SendPulse(pulseCounter, newPulses, modules, metadata)
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
		SendPulse(pulseCounter, []Pulse{{"button", "broadcaster", Low}}, modules, nil)
	}
	return pulseCounter.total()
}

type DestinationCycleData struct {
	targets           map[string]bool
	counters          map[string]int64
	expectedPulseType PulseType
	cycles            int64
}

func Part2(filePath string, finalDest string) int64 {
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

	/**
	 * Let's find the last conjunction module with more than one input
	 * Finding the cycle for these inputs should be enough to find the cycle for the whole system using the least common multiple
	 */
	finalModule, ok := modules[finalDest]
	if !ok {
		fmt.Println("finalDest module not found")
		return -1
	}

	var penultimateModule Module
	for moduleLabel := range finalModule.history {
		penultimateModule = modules[moduleLabel]
		break
	}

	metadataTargets := map[string]bool{}
	for moduleLabel := range penultimateModule.history {
		metadataTargets[moduleLabel] = true
	}

	// Find the conjunctions that need to be high before the final module can be high
	cycleData := DestinationCycleData{
		targets:           metadataTargets,
		counters:          make(map[string]int64),
		expectedPulseType: High,
		cycles:            1,
	}

	for {
		SendPulse(PulseCounter{}, []Pulse{{"button", "broadcaster", Low}}, modules, &cycleData)

		if len(cycleData.targets) == 0 {
			break
		}
		cycleData.cycles++
	}

	numbers := []int64{}
	for _, count := range cycleData.counters {
		numbers = append(numbers, count)
	}
	return findLCM(numbers)
}

func findLCM(numbers []int64) int64 {
	gcd := func(a, b int64) int64 { //general common divisor
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := func(a, b int64) int64 { //least common multiple
		return a * b / gcd(a, b)
	}

	result := int64(1)
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
}
