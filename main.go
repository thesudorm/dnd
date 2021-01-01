package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
	"math/rand"
	"time"
)

type Character struct {
	str int
	dex int
	con int
	int int
	wis int
	chr int
}

func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		PrintHelp()
		return
	}

	// someday might have to do some sort of fancy parsing in here for
	// command, input and flags
	command := os.Args[1]

	switch command {
		case "roll":
			RollDie(os.Args[2])
		default:
			fmt.Println(command, "is not a valid command.")
			PrintHelp()
	}
}

func GetInput() {

}

func RollDie(input string) {

	numOfDice, typeOfDice, err := ParseRollInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	results := []int{}
	total := 0

	for i := 0; i < numOfDice; i++ {
		randomInt := rand.Intn(typeOfDice) + 1
		results = append(results, randomInt)
		total += randomInt
	}

	fmt.Println(results)
	fmt.Println("Total:", total)
}

func ParseRollInput(input string) (int, int, error){
	parsed := strings.Split(input, "d")
	if parsed[0] == input {
		return -1, -1, errors.New("Not valid input for roll.")
	} else {
		numOfDice, err := strconv.Atoi(parsed[0])
		if err != nil {
			return -1, -1, errors.New("Invalid number of dice for roll command")
		}

		typeOfDice, err := strconv.Atoi(parsed[1])
		if err != nil {
			return -1, -1, errors.New("Invalid type of dice for roll command")
		}

		return numOfDice, typeOfDice, nil
	}
}

func PrintHelp() {
	fmt.Print(
	`Usage: dnd [COMMAND] [ARGUMENT]
	 Create and manage DnD 5e characters. Roll dice.

		roll [ARGUMENT] parses an argument to roll dice. i.e. 5d20, 1d6, etc
	`)
	fmt.Println()
}
