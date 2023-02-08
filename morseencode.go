package morsetranslator

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Encode a text message to a slice of morse sequences
func (t *MorseTranslator) EncodeMessageToMorse(input string) ([]Sequence, error) {

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

// Print morse sequence to text representation with separator
func (t *MorseTranslator) PrintRawSequences(sequences []Sequence) {
	for _, s := range sequences {
		print(t.SequencesTable[s].TextRepresentation)
	}
	fmt.Println()
}

// Print morse sequence to text representation without separator
func (t *MorseTranslator) PrintPrettySequences(Sequence []Sequence) {
	for _, s := range Sequence {
		if s == SHORT || s == LONG {
			print(t.SequencesTable[s].TextRepresentation)
		}

		if s == SEPARATOR_WORD {
			fmt.Print(" ")
		}
	}
}

func (t *MorseTranslator) Send(sequences []Sequence, sendfunc func(gpioPin string, state bool), print bool) {

	var len_seqs = len(sequences) - 1
	for i, s := range sequences {

		if print {
			if i == len_seqs {
				fmt.Print("fin")
				fmt.Printf("%s\n", t.SequencesTable[s].TextRepresentation)
			} else {
				fmt.Printf("%s", t.SequencesTable[s].TextRepresentation)
			}
		}

		switch s {
		case SHORT, LONG:
			sendfunc("18", true)
			time.Sleep(t.SequencesTable[s].Duration)
		case SEPARATOR, SEPARATOR_LETTER_IN_WORD, SEPARATOR_WORD:
			sendfunc("18", false)
			time.Sleep(t.SequencesTable[s].Duration)
		}
	}
}
