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

    // For now, load char sheet here by default. Can manage characters more effectively later
    pc := LoadSheet()

	// someday might have to do some sort of fancy parsing in here for
	// command, input and flags
	command, argument, err := GetCommandAndArgument()

	if err != nil {
		fmt.Println(err)
		PrintHelp()
		return;
	}

	// do some sort of input cleaning to make sure
	// we have an argument
	switch command {
		case "roll":
			numOfDice, typeOfDice, err := ParseRollInput(argument)

			if err != nil {
				fmt.Println(err)
			}

			results, total := RollDie(numOfDice, typeOfDice)
			fmt.Println(results)
		fmt.Println("Total:", total)
		case "create":
			// multiple steps
			// Pick Race
			// Pick Class
		case "check":
            AbilityCheck(pc, argument)
		case "char":
            // Print char sheet to 
		case "help":
            PrintHelp()
		default:
			fmt.Println(command, "is not a valid command.")
			PrintHelp()
	}
}

// For now, just two strings
func GetCommandAndArgument() (string, string, error){
	if len(os.Args) == 1 { // no command given, return error
		return "", "", errors.New("No command given.")
    } else if len(os.Args) == 2 {
        return os.Args[1], "", nil
    } else {
	    return os.Args[1], os.Args[2], nil
    }

}

func CreateCharacter() () {

}

func RollDie(numOfDice int, typeOfDice int) ([]int, int) {

	results := []int{}
	total := 0

	for i := 0; i < numOfDice; i++ {
		randomInt := rand.Intn(typeOfDice) + 1
		results = append(results, randomInt)
		total += randomInt
	}

	return results, total
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

		help    Print this help message
		roll    [ARGUMENT] Parses an argument to roll dice. i.e. 5d20, 1d6, etc
	`)
	fmt.Println()
}

func LoadSheet() (Character) {
    pc := Character {
        str: 18,
        dex: 15,
        con: 20,
        int: 13,
        wis: 14,
        chr: 10}

    return pc
}

func AbilityCheck(pc Character, ability string) () {
	results, total := RollDie(1, 20)

    toPrint, err := strconv.Atoi(total)

    switch ability {
        case "str":
        default:
            fmt.Println("Not an ability check")
    }
}

func Modifier(score int) int {
    modifier := 0
    switch {
        case score == 1:
            modifier = -5
        case score <= 3:
            modifier = -4
        case score <= 5:
            modifier = -3
        case score <= 7:
            modifier = -2
        case score <= 9:
            modifier = -1
        case score <= 11:
            modifier = 0
        case score <= 13:
            modifier = 1
        case score <= 15:
            modifier = 2
        case score <= 17:
            modifier = 3
        case score <= 19:
            modifier = 4
        case score <= 21:
            modifier = 5
        case score <= 23:
            modifier = 6
        case score <= 25:
            modifier = 7
    }
    return modifier
}
