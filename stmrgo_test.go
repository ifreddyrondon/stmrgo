package stmr_go_test

import (
	"testing"
	"github.com/ifreddyrondon/stmr-go"
)

func TestBasicStmr(t *testing.T)  {
	result := stmr_go.Stem("nationalization")
	if result != "nation" {
		t.Error("Expect nation. Got %v", result)
	}
}
