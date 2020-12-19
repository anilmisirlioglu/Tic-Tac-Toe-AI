package game

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	old := os.Stdout // backup real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Write("test context", NextLineBehind, true, 2)

	outChannel := make(chan string)
	go func() {
		var buffer bytes.Buffer
		_, _ = io.Copy(&buffer, r)
		outChannel <- buffer.String()
	}()

	_ = w.Close()
	os.Stdout = old
	text := <-outChannel

	expected := "[*] test context\n\n"
	if text != expected {
		t.Errorf("Actual: %s, Expected: %s", text, expected)
	}
}
