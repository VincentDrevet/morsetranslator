package morsetranslator

import (
	"errors"
	"fmt"
	"strings"
)

func EncodeMessageToMorse(input string) ([]Sequence, error) {

	//Split message by " "
	split := strings.Split(input, " ")
	nwords := len(split) - 1

	finalSequence := []Sequence{}

	for i, word := range split {

		ncharacter := len(word) - 1

		// Processing word
		for j, c := range word {

			if letter, ok := LETTER[c]; ok {
				finalSequence = append(finalSequence, letter...)
			} else {
				return []Sequence{}, errors.New(fmt.Sprintf("unknown letter at word %d character %d", i+1, j+1))
			}
			if !(ncharacter == j) {
				finalSequence = append(finalSequence, SEPARATOR_LETTER_IN_WORD)
			}
		}
		if !(nwords == i) {
			finalSequence = append(finalSequence, SEPARATOR_WORD)
		}
	}

	return finalSequence, nil
}

func PrintRawSequences(sequences []Sequence) {
	for _, s := range sequences {
		s.String()
	}
	fmt.Println()
}

func PrintPrettySequences(Sequence []Sequence) {
	for _, s := range Sequence {
		if s == SHORT || s == LONG {
			s.String()
		}

		if s == SEPARATOR_WORD {
			fmt.Print(" ")
		}
	}
}

func DecodeMessageToMorse()
