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
	value        int
	nextId       string
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
			if part.variables[variableName] < rule.value {
				return rule.nextId
			}
		} else {
			if part.variables[variableName] > rule.value {
				return rule.nextId
			}
		}
	}
	return ruleSet.fallbackRule
}

func Part1(filePath string) int {
	lines := aoc.ReadFileToLines(filePath)
	ruleRegex := regexp.MustCompile(`^[a-z]+\{`)
	comparableRegex := regexp.MustCompile(`^([a-z]+)([<>])([0-9]+)$`)

	parts := make([]Part, 0)
	ruleSets := make([]RuleSet, 0)

	for _, line := range lines {
		if ruleRegex.MatchString(line) {
			// get value before first "{"
			index := strings.Index(line, "{")
			ruleSet := RuleSet{id: line[:index]}
			ruleSet.rules = make([]Rule, 0)

			// get value after first "{"
			line = line[index+1:]
			// get value before last "}"
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
					// get value before ":"
					comparableString := strings.TrimSpace(ruleParts[0])
					comparableRegexParts := comparableRegex.FindStringSubmatch(comparableString)
					rule.variableName = comparableRegexParts[1]
					rule.comparable = comparableRegexParts[2]
					rule.value, _ = strconv.Atoi(comparableRegexParts[3])
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

		fmt.Printf("Part: %v, Result: %s\n", part, next)
	}

	return total
}

func main() {
	result := Part1("test input")
	fmt.Println(result)
}
