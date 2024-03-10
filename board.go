package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/gregoryv/english"
)

const win, lost, playing = "WIN", "LOST", "PLAYING"

type Board struct {
	guesses, corrects, wrongs []string
	secretWord                string
	progress                  string
	gameStatus                string
}

func NewBoard() *Board {
	p := new(Board)
	p.secretWord = strings.ToUpper(english.RandomWord())
	p.gameStatus = playing
	p.progress = regexp.MustCompile(`[a-zA-Z]`).ReplaceAllString(p.secretWord, "_")
	return p
}

func (thiss *Board) ShowBoard() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------Guess Me--------------")
	fmt.Println("")
	thiss.printProgress()
	thiss.printHangMan()
    thiss.printWrongs()
	// thiss.revealCorrectAnswer()
	if thiss.gameStatus == playing {
		thiss.makeAGuess()
	} else {
        if thiss.gameStatus == win {
            fmt.Println("Congratulations")
        }
        if thiss.gameStatus == lost {
            fmt.Println("Answer was:")
            thiss.revealCorrectAnswer()
        }
        if thiss.gameStatus == lost || thiss.gameStatus == win  {
            fmt.Println("")
            fmt.Println("")
            fmt.Println("--------------")
            fmt.Println("Starting New Game....")
        }
    }
}

func (thiss *Board) makeAGuess() {
	guess := expectALetter()
	validate := isLetter(guess)
	if validate {
		thiss.checkAnswer(guess)
		thiss.checkScore()
	}
	thiss.ShowBoard()
}

func (thiss *Board) GetState() string {
	return thiss.gameStatus
}

func (thiss *Board) checkAnswer(guess string) {
	guess = strings.ToUpper(strings.TrimSpace(guess))


    if slices.Contains(thiss.guesses, guess){
        return
    }

    thiss.guesses = append(thiss.guesses, guess)

	if strings.Contains(thiss.secretWord, guess) {
		thiss.corrects = append(thiss.corrects, guess)
		thiss.replaceCorrectLetters(guess)
	} else {
		thiss.wrongs = append(thiss.wrongs, guess)
	}
}

func (thiss *Board) checkScore() {

	if len(thiss.wrongs) == 6 {
		thiss.gameStatus = lost
	}

	if !strings.Contains(thiss.progress, "_") {
		thiss.gameStatus = win
	}
}
func (thiss *Board) printProgress() {
    fmt.Printf("??? %s ??? (%d letters)", thiss.progress, len(thiss.secretWord)) 
}

func (thiss *Board) printWrongs() {
    if len(thiss.wrongs) > 0 {
        fmt.Printf("Wrong guesses: %s ", thiss.wrongs) 
    }
    fmt.Println("")
    fmt.Println("")
}

func (thiss *Board) replaceCorrectLetters(guess string) {
	progressByteArray := []byte(thiss.progress)
	secretAnswer := thiss.secretWord

	for i, runeVal := range secretAnswer {
		letter := string(runeVal)
		if letter == guess {
			progressByteArray[i] = secretAnswer[i]
		}

	}
	thiss.progress = string(progressByteArray)
}

func (thiss *Board) revealCorrectAnswer() {
	fmt.Println(thiss.secretWord)
}

func isLetter(input string) bool {
	retVal := false
	retVal = len(strings.TrimSpace(input)) == 1 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(input)
	if !retVal {
		fmt.Println("Please enter only one letter. Must be a-Z")
	}
	return retVal
}

func expectALetter() string {
	fmt.Print("Guess a letter: ")
	var retVal string
	fmt.Scanln(&retVal)
	return retVal
}

func (thiss *Board) printHangMan() {
	result := [7]string{
		`

                     +---+
                         |
                         |
                         |
                        ===
`,
		`

                     +---+
                         |
                     |   |
                         |
                        ===
`,
		`

                     +---+
                         |
                    /|   |
                         |
                        ===
`,
		`

                     +---+
                         |
                    /|\  |
                         |
                        === 
`,
		`

                     +---+
                         |
                    /|\  |
                    /    |
                        === 
`,
		`

                     +---+
                         |
                    /|\  |
                    / \  |
                        === 
`,
		`

                     +---+
                     0   |
                    /|\  |
                    / \  |
                        === 
                    Game Over!    
`,
	}

	fmt.Print(result[len(thiss.wrongs)])
}
