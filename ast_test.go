package lang

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	p := Program{
		scope: &Scope{
			values: make(map[string]*Literal),
		},
	}

	v, err := NewLiteral("raw")
	if err != nil {
		t.Error("invalid value")
	}

	assign := &AssignmentStatement{
		Name:  "potato",
		Type:  "string",
		Value: v,
	}

	p.statements = append(p.statements, assign)

	err = p.Run()
	if err != nil {
		t.Error(err)
	}

	if str := p.scope.Get("potato").MustString(); str != "raw" {
		fmt.Println("Did not get correct value. Got", str, "expected", "raw")
	}
}
