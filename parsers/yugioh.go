package parsers

import "github.com/alecthomas/participle/v2/lexer"

type YuGiOhDescription struct {
	Effects []YuGiOhEffect `parser:"@@*"`
}

type YuGiOhEffect struct {
	Condition []YuGiOhEffectCondition `parser:"@@*"`
	Statement string                  `parser:"| @Ident*"`
}

type YuGiOhEffectCondition struct {
	Pos lexer.Position

	Subject  YuGiOhSubject `parser:"'If' @@"`
	Operator bool          `parser:"@'is'"`
	State    string        `parser:"@Ident*(':'? | ','? )!"`
}

type YuGiOhSubject struct {
	Pos lexer.Position

	Self  bool  `parser:"'this' 'card'"`
	Value Value `parser:"| @@"`
}

type Value struct {
	Pos lexer.Position

	Number *int    `parser:"@Int"`
	String *string `parser:"| @Ident"`
}
