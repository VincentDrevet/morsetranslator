package morsetranslator

import "time"

type MorseTranslator struct {
	Configuration  Configuration
	SequencesTable map[Sequence]SequenceSettings
}

type Configuration struct {
	ShortDuration                           time.Duration `json:"short_duration"`
	ShortTextRepresentation                 string        `json:"short_text_representation"`
	LongTextRepresentation                  string        `json:"long_text_representation"`
	SeparatorTextRepresentation             string        `json:"separator_text_representation"`
	SeparatorLetterInWordTextRepresentation string        `json:"separator_letter_in_word_text_representation"`
	SeparatorWordTextRepresentation         string        `json:"separator_word_text_representation"`
}

func NewMorseTranslator(configuration Configuration) MorseTranslator {
	var sequencesTable = map[Sequence]SequenceSettings{
		SHORT: SequenceSettings{
			Duration:           configuration.ShortDuration,
			TextRepresentation: configuration.ShortTextRepresentation,
		},

		LONG: SequenceSettings{
			Duration:           configuration.ShortDuration * 3,
			TextRepresentation: configuration.LongTextRepresentation,
		},

		SEPARATOR: SequenceSettings{
			Duration:           configuration.ShortDuration,
			TextRepresentation: configuration.SeparatorTextRepresentation,
		},

		SEPARATOR_LETTER_IN_WORD: SequenceSettings{
			Duration:           configuration.ShortDuration * 3,
			TextRepresentation: configuration.SeparatorLetterInWordTextRepresentation,
		},

		SEPARATOR_WORD: SequenceSettings{
			Duration:           7 * (configuration.ShortDuration * 3),
			TextRepresentation: configuration.SeparatorWordTextRepresentation,
		},
	}

	return MorseTranslator{
		Configuration:  configuration,
		SequencesTable: sequencesTable,
	}

}
