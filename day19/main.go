package main

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	variableName string
	comparable   string
	quantity     int
	nextId       string
}

type RuleType int

const (
	ACCEPTED RuleType = iota
	REJECTED
	FOLLOW
)

func (rule *Rule) ruleType() RuleType {
	if rule.nextId == "A" {
		return ACCEPTED
	} else if rule.nextId == "R" {
		return REJECTED
	} else {
		return FOLLOW
	}
}

func (ruleSet *RuleSet) fallbackRuleType() RuleType {
	if ruleSet.fallbackRule == "A" {
		return ACCEPTED
	} else if ruleSet.fallbackRule == "R" {
		return REJECTED
	} else {
		return FOLLOW
	}
}

type RuleSet struct {
	id           string
	rules        []Rule
	fallbackRule string
}

type Part struct {
	variables map[string]int
}

func (part Part) getTotal() int {
	total := 0
	for _, value := range part.variables {
		total += value
	}
	return total
}

func (ruleSet *RuleSet) findNext(part Part) string {
	for _, rule := range ruleSet.rules {
		variableName := rule.variableName
		if rule.comparable == "<" {
			if part.variables[variableName] < rule.quantity {
				return rule.nextId
			}
		} else {
			if part.variables[variableName] > rule.quantity {
				return rule.nextId
			}
		}
	}
	return ruleSet.fallbackRule
}

func parseInput(filePath string) ([]Part, []RuleSet) {
	lines := aoc.ReadFileToLines(filePath)
	ruleRegex := regexp.MustCompile(`^[a-z]+\{`)
	comparableRegex := regexp.MustCompile(`^([a-z]+)([<>])([0-9]+)$`)

	parts := make([]Part, 0)
	ruleSets := make([]RuleSet, 0)

	for _, line := range lines {
		if ruleRegex.MatchString(line) {
			// get quantity before first "{"
			index := strings.Index(line, "{")
			ruleSet := RuleSet{id: line[:index]}
			ruleSet.rules = make([]Rule, 0)

			// get quantity after first "{"
			line = line[index+1:]
			// get quantity before last "}"
			index = strings.LastIndex(line, "}")
			line = line[:index]
			// split on ","
			ruleStrings := strings.Split(line, ",")
			for _, ruleString := range ruleStrings {
				rule := Rule{}
				ruleString = strings.TrimSpace(ruleString)
				// split on ":"
				ruleParts := strings.Split(ruleString, ":")

				if len(ruleParts) == 1 {
					// no ":" found, so this is the fallback rule
					ruleSet.fallbackRule = strings.TrimSpace(ruleParts[0])
					continue
				} else {
					// get quantity before ":"
					comparableString := strings.TrimSpace(ruleParts[0])
					comparableRegexParts := comparableRegex.FindStringSubmatch(comparableString)
					rule.variableName = comparableRegexParts[1]
					rule.comparable = comparableRegexParts[2]
					rule.quantity, _ = strconv.Atoi(comparableRegexParts[3])
					rule.nextId = strings.TrimSpace(ruleParts[1])
					ruleSet.rules = append(ruleSet.rules, rule)
				}
			}
			ruleSets = append(ruleSets, ruleSet)
		} else if line != "" {
			// trim {}
			line = line[1 : len(line)-1]
			// splits on ","
			partStrings := strings.Split(line, ",")
			part := Part{variables: make(map[string]int)}
			for _, partString := range partStrings {
				partString = strings.TrimSpace(partString)
				// split on "="
				partStringParts := strings.Split(partString, "=")
				variableName := strings.TrimSpace(partStringParts[0])
				variableValue, _ := strconv.Atoi(strings.TrimSpace(partStringParts[1]))
				part.variables[variableName] = variableValue
			}
			parts = append(parts, part)
		}
	}

	return parts, ruleSets
}

func Part1(filePath string) int {
	parts, ruleSets := parseInput(filePath)

	ruleSetMap := make(map[string]RuleSet)
	for _, ruleSet := range ruleSets {
		ruleSetMap[ruleSet.id] = ruleSet
	}

	total := 0
	for _, part := range parts {
		next := "in"
		done := false
		for !done {
			ruleSet, _ := ruleSetMap[next]
			next = ruleSet.findNext(part)
			if next == "A" || next == "R" {
				done = true
			}
		}

		if next == "A" {
			total += part.getTotal()
		}

		// fmt.Printf("Part: %v, Result: %s\n", part, next)
	}

	return total
}

type Range struct {
	start, end int
}

func (r Range) possibleValues() int {
	return r.end - r.start + 1
}

var variablesMap = map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}

func doCombinations(ruleSetMap map[string]RuleSet, rule string, inputRanges []Range) int64 {
	fmt.Println(rule, "Ranges: ", inputRanges)
	result := int64(0)
	resultSet := ruleSetMap[rule]

	// for each rule in the ruleset
	for _, nextRule := range resultSet.rules {
		newRanges := make([]Range, len(inputRanges))
		// copy the inputRanges
		copy(newRanges, inputRanges)

		// get the index of the variable
		variableId := variablesMap[nextRule.variableName]

		if nextRule.comparable == "<" {
			newRanges[variableId].end = nextRule.quantity - 1 // These values are allowed
			inputRanges[variableId].start = nextRule.quantity // The rest go to the other rules
			switch nextRule.ruleType() {
			case FOLLOW:
				// Pass the new ranges to the next rule
				result += doCombinations(ruleSetMap, nextRule.nextId, newRanges)
			case ACCEPTED:
				// Calculate the possible combinations for the new ranges
				result += doRangeProduct(newRanges)
			case REJECTED:
			}
		} else if nextRule.comparable == ">" {
			newRanges[variableId].start = nextRule.quantity + 1 // These values are allowed
			inputRanges[variableId].end = nextRule.quantity     // The rest go to the other rules
			switch nextRule.ruleType() {
			case FOLLOW:
				// Pass the new ranges to the next rule
				result += doCombinations(ruleSetMap, nextRule.nextId, newRanges)
			case ACCEPTED:
				// Calculate the possible combinations for the new ranges
				result += doRangeProduct(newRanges)
			case REJECTED:
			}
		}
	}

	// Run any remaining ranges through the fallback rule
	switch resultSet.fallbackRuleType() {
	case FOLLOW:
		result += doCombinations(ruleSetMap, resultSet.fallbackRule, inputRanges)
	case ACCEPTED:
		result += doRangeProduct(inputRanges)
	case REJECTED:
	}

	return result
}

func doRangeProduct(ranges []Range) int64 {
	result := 1
	fmt.Println("Range product : Ranges", ranges)
	for _, r := range ranges {
		result *= r.possibleValues()
		fmt.Println("Result:", result)
	}
	fmt.Println("Result:", result)
	return int64(result)
}

func Part2(filePath string) int64 {
	_, ruleSets := parseInput(filePath)

	// create a map of the rulesets
	ruleSetMap := map[string]RuleSet{}
	for _, ruleSet := range ruleSets {
		ruleSetMap[ruleSet.id] = ruleSet
	}

	initialRanges := []Range{
		{1, 4000},
		{1, 4000},
		{1, 4000},
		{1, 4000},
	}
	return doCombinations(ruleSetMap, "in", initialRanges)
}

func main() {
	result := Part1("test input")
	fmt.Println(result)
}
