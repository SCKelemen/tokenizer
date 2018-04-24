package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Running")
	filepath.Walk("C:\\Users\\skelemen\\Desktop\\tokentest", walk)
	fmt.Println("done.")
}

func walk(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	tokenizer := TextTokenizer{}
	if !tokenizer.canTokenizeDocument(path) {
		fmt.Println(path)
		return nil
	}
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	scanner := Scanner{file: file}
	fmt.Printf("Scanning %s...\n", path)
	scanner.Scan()
	return nil
}

type TextTokenizer struct {
}

/*
func (t TextTokenizer) Tokenize(path string) []string {
	if !t.canTokenizeDocument(path) {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	buffer := make([]byte, 3)
	n1, err := f.Read(buffer)
}
*/
func (t TextTokenizer) canTokenizeDocument(path string) bool {
	extension := filepath.Ext(path)
	switch extension {
	case ".text":
	case ".txt":
		fmt.Println("txt")
		return true
	default:
		fmt.Println("not text")
		return false
	}
	return false
}

func (s Scanner) Emit(token string) {
	fmt.Printf("\tTOKEN:\t%s\n", token)
}

type Scanner struct {
	file   *os.File
	buffer [3]byte
}

func (s Scanner) Scan() {
	b1 := make([]byte, 3)
	count, err := s.file.Read(b1)
	if err == io.EOF {
		return
	}
	ibuf := make([]byte, 3)
	buf := bytes.NewBuffer(ibuf)
	switch count {
	case 1:
		buf.WriteByte(b1[0])
		buf.WriteRune('\x00')
		buf.WriteRune('\x00')
	case 2:
		buf.WriteByte(b1[0])
		buf.WriteByte(b1[1])
		buf.WriteRune('\x00')
	case 3:
		buf.WriteString(string(b1))
	default:
		return
	}
	s.Emit(buf.String())
	s.Scan()
}
