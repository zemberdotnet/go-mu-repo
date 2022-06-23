package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type OutputWriter struct {
	buf      *bytes.Buffer
	target   string
	exitCode int
}

func NewOuputWriter(target string) *OutputWriter {
	return &OutputWriter{
		buf:      &bytes.Buffer{},
		target:   target,
		exitCode: 0,
	}
}

func (o *OutputWriter) Write(p []byte) (n int, err error) {
	return o.buf.Write(p)
}

func (o *OutputWriter) Flush() {
	bytes := o.buf.Bytes()
	s := struct {
		Target   string `json:"target"`
		Output   string `json:"output"`
		ExitCode int    `json:"ExitCode"`
	}{
		Target:   o.target,
		Output:   string(bytes),
		ExitCode: o.exitCode,
	}

	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
	// TODO return string instead of printing

}
