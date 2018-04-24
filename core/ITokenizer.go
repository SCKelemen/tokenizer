package core

type ITokenizer interface {
	Tokenize(path string) []string
}

type Token string
