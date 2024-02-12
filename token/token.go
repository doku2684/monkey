package token


type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers and literals
  IDENT = "IDENT" // add, foobar, x, y
  INT = "INT" // 153245
  STRING = "STRING"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"

  LT = "<"
  GT = ">"

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"

  // Two character token
  EQ = "=="
  NOT_EQ = "!="
)

var keywords = map[string] TokenType {
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
