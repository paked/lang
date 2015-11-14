package lang

type Scope struct {
	values map[string]string
}

func (s *Scope) Set(key, val string) {
	s.values[key] = val
}

func (s *Scope) Get(key string) string {
	return s.values[key]
}

type Program struct {
	scope      *Scope
	statements []Statement
}

func (prog *Program) Run() error {
	for _, stmt := range prog.statements {
		stmt.Eval(prog.scope)
	}

	return nil
}

type Statement interface {
	Eval(*Scope) error
}

type AssignmentStatement struct {
	Name  string
	Type  string
	Value string
}

func (as *AssignmentStatement) Eval(s *Scope) error {
	s.Set(as.Name, as.Value)

	return nil
}
