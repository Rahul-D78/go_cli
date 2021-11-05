package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Hello World")
	dice := flag.String("d", "d6", "The type of dice to role the default would be :d6")
	numRoll := flag.Int("n", 1, "The number of times to roll the dice")
	sum := flag.Bool("s", false, "Print the sum of the dice")
	advantage := flag.Bool("a", false, "Roll with advantage")
	disadvantage := flag.Bool("d", false, "Roll with disadvantage")
	flag.Parse()
	// fmt.Printf("You choose a %s\n", *dice)

	//type casting to byte slice
	matched, _ := regexp.Match("d\\d+", []byte(*dice))
	// fmt.Println(matched)

	if matched == true {
		//1 to the end of the string
		rolls := rollDice(dice, numRoll)
		printDice(rolls)

		if *sum == true {
			diceSum := sumDice(rolls)
			fmt.Printf("The sum of the dice is %d\n", diceSum)
		}
		if *advantage == true {
			diceSum := rollWithAdvantage(rolls)
			fmt.Printf("The sum of the dice is %d\n", diceSum)
		}
		if *disadvantage == true {
			diceSum := disRollWithAdvantage(rolls)
			fmt.Printf("The sum of the dice is %d\n", diceSum)
		}
	} else {
		log.Fatalf("Invalid dice type")
	}
}

func rollDice(dice *string, times *int) []int {
	var rolls []int
	diceSides := (*dice)[1:]
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}
	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, dice := range rolls {
		sum += dice
	}
	return sum
}

func rollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func disRollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}
