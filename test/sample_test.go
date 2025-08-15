package test

import (
	"testing"

	"github.com/yourusername/yal/internal/interpreter"
)

func TestSampleProgram(t *testing.T) {
	interp := interpreter.New()
	err := interp.Run("print('Test')")
	if err != nil {
		t.Fatalf("Interpreter failed: %v", err)
	}
}
