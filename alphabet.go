package morsetranslator

import (
	"time"
)

type Sequence uint8

const (
	SHORT                    = 1 // ti .
	LONG                     = 2 // taah _
	SEPARATOR                = 3 // Separator between . and _ in a letter (.)
	SEPARATOR_LETTER_IN_WORD = 4 // Serapator between letter in a word (_)
	SEPARATOR_WORD           = 5 // Separator between word (_______)
)

type SequenceSettings struct {
	Duration           time.Duration
	TextRepresentation string
}

var LETTER = map[rune][]Sequence{
	'a':  []Sequence{SHORT, SEPARATOR, LONG},                                                                        // a = ._
	'b':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},                                    // b = _...
	'c':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},                                     // c = _._.
	'd':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT},                                                      // d = _..
	'e':  []Sequence{SHORT},                                                                                         // e = .
	'f':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},                                    //f = .._.
	'g':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, SHORT},                                                       //g = __.
	'h':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},                                   //h = ....
	'i':  []Sequence{SHORT, SEPARATOR, SHORT},                                                                       //i = ..
	'j':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG},                                      //j = .___
	'k':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, LONG},                                                       //k = _._
	'l':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT},                                    //l = ._..
	'm':  []Sequence{LONG, SEPARATOR, LONG},                                                                         //m = __
	'n':  []Sequence{LONG, SEPARATOR, SHORT},                                                                        //n = _.
	'o':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, LONG},                                                        //o = ___
	'p':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT},                                     //p =.__.
	'q':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, LONG},                                      //q = __._
	'r':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},                                                      //r =._.
	's':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},                                                     //s = ...
	't':  []Sequence{LONG},                                                                                          //t = _
	'u':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG},                                     //u = ..__
	'v':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG},                                    //v = ..._
	'w':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, LONG},                                                       //w = .__
	'x':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG},                                     //x = _.._
	'y':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG},                                      //y = _.__
	'z':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT},                                     //z = __..
	'0':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG},                      //0 = _____
	'1':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG},                     //1 = .____
	'2':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG},                    //2 = ..___
	'3':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG},                   //3 = ...__
	'4':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG},                  //4 = ...._
	'5':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},                 //5 = .....
	'6':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},                  //6 = _....
	'7':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},                   //7 = __...
	'8':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT},                    //8 = ___..
	'9':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT},                     //9 = ____.
	',':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG},   //, = __..__
	'?':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT}, //? = ..__..
	':':  []Sequence{LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT},  //: = ___...
	'-':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG}, //- = _...._
	'"':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT}, // " = ._.._.
	'(':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT},                    // ( = _.__.
	'=':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG},                   // = = _..._
	';':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},  // ; = _._._.
	'/':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},                   // / = _.._.
	'\'': []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT},   // ' = .____.
	'_':  []Sequence{SHORT, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, LONG},  // _ = ..__._
	')':  []Sequence{LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, LONG},   // ) = _.__._
	'+':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},                   // + = ._._.
	'@':  []Sequence{SHORT, SEPARATOR, LONG, SEPARATOR, LONG, SEPARATOR, SHORT, SEPARATOR, LONG, SEPARATOR, SHORT},  //@ = .__._.
}
