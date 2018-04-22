package bytecounter

import (
	"bufio"
	"strings"
)

type Bytecounter int

func (c *Bytecounter) Write(p []byte) (int, error) {
	*c += Bytecounter(len(p))
	return len(p), nil
}

func count(p []byte, splitFunc bufio.SplitFunc) int {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(splitFunc)

	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c = WordCounter(count(p, bufio.ScanWords))
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	*c = LineCounter(WordCounter(count(p, bufio.ScanLines)))
	return len(p), nil
}
