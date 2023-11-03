package testing

import (
	"fmt"
)

type FakeTLogger struct {
	Messages []string
}

func (tl *FakeTLogger) Log(args ...interface{}) {
	for _, arg := range args {
		tl.Messages = append(tl.Messages, fmt.Sprint(arg))
	}
}

func (tl *FakeTLogger) Helper() {}
