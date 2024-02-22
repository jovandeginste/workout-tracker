package ast

type Token string

const (
	eof            Token = "eof"
	whitespace     Token = "ws"
	failure        Token = "failure"
	Value          Token = "Value" // 1, 2, 100, etc.
	Operand        Token = "n"
	plural         Token = "plural="
	nPlurals       Token = "nplurals="
	Equal          Token = "=="
	assign         Token = "="
	Greater        Token = ">"
	GreaterOrEqual Token = ">="
	Less           Token = "<"
	LessOrEqual    Token = "<="
	Reminder       Token = "%"
	NotEqual       Token = "!="
	LogicalAnd     Token = "&&"
	LogicalOr      Token = "||"
	Question       Token = "?"
	colon          Token = ":"
	semicolon      Token = ";"
	leftBracket    Token = "("
	rightBracket   Token = ")"
)
