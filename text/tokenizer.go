package text

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

type TextTokenizer struct {
}

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

func (t TextTokenizer) canTokenizeDocument(path string) bool {
	extension := filepath.Ext(path)
	switch extension {
	case "text":
	case "txt":
		return true
	default:
		return false
	}
	return false
}

func (t TextTokenizer) Emit(token string) {

}

type Scanner struct {
	file   os.File
	buffer [3]byte
}

func (s Scanner) Scan() {
	b1 := make([]byte, 3)
	count, err := s.file.Read(b1)
	if err == io.EOF {
		return nil
	}
	buf := bytes.Buffer(3)
	switch count {
	case 1:
		buf.WriteRune(b1[0])
		buf.WriteRune('\x00')
		buf.WriteRune('\x00')
	case 2:
		buf.WriteRune(b1[0])
		buf.WriteRune(b1[1])
		buf.WriteRune('\x00')
	case 3:
		buf.WriteString(string(b1))
	default: 
		return nil
	}
	s.Emit(buf)
	s.Scan()
}

func (s Scanner) cleanAndEmit(data []byte) {
	switch 
}