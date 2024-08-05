package main

import (
	"bufio"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Board - the game board
type Board struct {
	colorsInPlay []string
	pegsInPlay   int
	solution     []string
}

func main() {
	currentGame := setBoard()
	gameInPlay := true

	for index := 1; gameInPlay; index++ {
		printBoard(currentGame)
		fmt.Println("What is your guess?  Please separate your colors with a comma")
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if guess[0:4] == "quit" {
			break
		}

		noSpaceGuess := strings.Replace(guess, " ", "", -1)
		if reflect.DeepEqual(currentGame.solution, strings.Split(noSpaceGuess[:len(noSpaceGuess)-2], ",")) {
			fmt.Println("YOU WIN!! And it only took you ", index, " guesses!")
			gameInPlay = false
		} else {
			checkGuess(noSpaceGuess, currentGame)
			fmt.Println("Please try again")
		}
	}
}

func setBoard() (currentGame Board) {
	fmt.Println("Hello Player 1")
	fmt.Println("How many pegs? Please pick between 2 and 20")
	reader := bufio.NewReader(os.Stdin)
	numPegs, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		currentGame.pegsInPlay, _ = strconv.Atoi(numPegs[:len(numPegs)-2])
		if currentGame.pegsInPlay > 1 && currentGame.pegsInPlay < 21 {
			fmt.Println("Yay! You picked ", currentGame.pegsInPlay, " pegs")
		} else {
			fmt.Println(currentGame.pegsInPlay, "? I do not think so.  I am thinking 4 sounds good.")
			currentGame.pegsInPlay = 4
		}
	}

	var intColors int
	fmt.Println("How many colors?  Please pick between 2 and 20")
	numColors, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		intColors, _ = strconv.Atoi(numColors[:len(numColors)-2])
		if intColors > 1 && intColors < 21 {
			fmt.Println("Yay! You picked ", intColors, " colors")
		} else {
			fmt.Println(intColors, "? I do not think so.  I am thinking 4 sounds good.")
			intColors = 4
		}
	}
	currentGame.colorsInPlay = getColors(intColors)
	var b [8]byte
	_, err = crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	for index := 0; index < currentGame.pegsInPlay; index++ {
		currentGame.solution = append(currentGame.solution, currentGame.colorsInPlay[rand.Int()%intColors])
	}
	return
}

func printBoard(currentGame Board) {
	fmt.Println("The possible colors are", currentGame.colorsInPlay)
}

func getColors(numColors int) (colors []string) {
	palette := []string{"blue", "red", "yellow", "green", "white", "pink", "brown", "purple", "black", "orange", "gold", "silver", "copper", "rose", "teal", "salmon", "olive", "aquamarine", "amethyst", "cyan"}
	colors = palette[0:numColors]
	return
}

func checkGuess(guess string, currentGame Board) {
	var correctPlace, correctColor int
	guessSlice := strings.Split(guess[:len(guess)-2], ",")
	for i, peg := range currentGame.solution {
		if guessSlice[i] == peg {
			correctPlace = +1
		}
	}

	fmt.Println("You have correctly guessed ", correctPlace, " pegs")
	fmt.Println("You have also guessed ", correctColor, " colors")
}
