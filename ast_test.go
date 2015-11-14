package lang

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	p := Program{
		scope: &Scope{
			values: make(map[string]string),
		},
	}

	assign := &AssignmentStatement{
		Name:  "potato",
		Type:  "string",
		Value: "raw",
	}

	p.statements = append(p.statements, assign)

	err := p.Run()
	if err != nil {
		t.Error(err)
	}

	if str := p.scope.Get("potato"); str != "raw" {
		fmt.Println("Did not get correct value. Got", str, "expected", "raw")
	}
}
