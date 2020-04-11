package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		NewLog()
		Info("hello", "world")
		Infof("hello %s", "world")
	})
}
