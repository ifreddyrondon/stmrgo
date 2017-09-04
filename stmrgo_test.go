package stmrgo_test

import (
	"testing"
	"github.com/ifreddyrondon/stmrgo"
)

func TestBasicStmr(t *testing.T)  {
	result := stmrgo.Stem("nationalization")
	if result != "nation" {
		t.Error("Expect nation. Got %v", result)
	}
}
