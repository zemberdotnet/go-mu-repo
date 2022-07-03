package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// OutputWriterCollection holds and manages a collection of OutputWriters
// It is thread safe and can construct output from many OutputWriters
type OutputWriterCollection struct {
	writers []OutputWriter
	mutex   *sync.Mutex
}

func NewOutputWriterCollection() *OutputWriterCollection {
	return &OutputWriterCollection{
		writers: make([]OutputWriter, 0),
		mutex:   &sync.Mutex{},
	}
}

// Add adds a OutputWriter to the OutputWriter collection
func (owc *OutputWriterCollection) Add(ow OutputWriter) {
	owc.mutex.Lock()
	owc.writers = append(owc.writers, ow)
	owc.mutex.Unlock()
}

// Flush flushes the OutputWriters to Stdout
func (owc *OutputWriterCollection) Flush() error {

	if JsonFlag {
		output := make([]json.RawMessage, 0)
		for _, ow := range owc.writers {
			flushed := json.RawMessage(ow.Flush())
			output = append(output, flushed)
		}
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(output)
	} else {
		output := make([]byte, 0)
		for _, ow := range owc.writers {
			flushed := ow.Flush()
			output = append(output, flushed...)
		}
		fmt.Fprint(os.Stdout, string(output))
		return nil
	}

}

type OutputWriter interface {
	Flush() []byte
	io.Writer
}

func CreateOutputWriter(target string) OutputWriter {
	if JsonFlag {
		return NewJsonOutputWriter(target)
	} else {
		return NewStdOutputWriter(target)
	}
}

var _ OutputWriter = &JsonOutputWriter{}
var _ OutputWriter = &StdOutputWriter{}

type StdOutputWriter struct {
	buf    *bytes.Buffer
	target string
}

func NewStdOutputWriter(target string) *StdOutputWriter {
	return &StdOutputWriter{
		buf:    &bytes.Buffer{},
		target: target,
	}
}

// Write implements the io.Writer interface
func (o *StdOutputWriter) Write(p []byte) (n int, err error) {
	return o.buf.Write(p)
}

// Flush flushes the OutputWriter to a json.RawMessage
func (o *StdOutputWriter) Flush() []byte {
	prefix := fmt.Sprintf("#target:%v\n", o.target)
	return append([]byte(prefix), o.buf.Bytes()...)
}

type JsonOutputWriter struct {
	buf      *bytes.Buffer
	target   string
	exitCode int
}

// NewJsonOutputWriter creates a new OutputWriter generating JSON output
// from a Command
func NewJsonOutputWriter(target string) *JsonOutputWriter {
	return &JsonOutputWriter{
		buf:      &bytes.Buffer{},
		target:   target,
		exitCode: 0,
	}
}

// Write implements the io.Writer interface
func (o *JsonOutputWriter) Write(p []byte) (n int, err error) {
	return o.buf.Write(p)
}

// Flush flushes the OutputWriter to a json.RawMessage
func (o *JsonOutputWriter) Flush() []byte {
	bytes := o.buf.Bytes()
	c := map[string]interface{}{}
	// We need to determine if the output is a JSON object or just bytes
	err := json.Unmarshal(bytes, &c)
	if err == nil {
		s := struct {
			Target   string          `json:"target"`
			Output   json.RawMessage `json:"output"`
			ExitCode int             `json:"exitCode"`
		}{
			Target:   o.target,
			Output:   json.RawMessage(bytes),
			ExitCode: o.exitCode,
		}
		buf, err := json.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		return buf

	} else {
		s := struct {
			Target   string `json:"target"`
			Output   string `json:"output"`
			ExitCode int    `json:"exitCode"`
		}{
			Target:   o.target,
			Output:   string(bytes),
			ExitCode: o.exitCode,
		}

		buf, err := json.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		return buf
	}
}
