package gokatas

type EvalError struct {
	text string
}

func (e EvalError) Error() string {
	return e.text
}

type Env map[string]int

type Node interface {
	eval(e Env) (int, error)
}

type Litt struct {
	value int
}

func (t Litt) eval(e Env) (int, error) {
	return t.value, nil
}

type Var struct {
	name string
}

func (t Var) eval(e Env) (int, error) {
	value, ok := e[t.name]
	if ok {
		return value, nil
	} else {
		return 0.0, EvalError{text: "Unknown variable " + t.name}
	}
}

type Plus struct {
	left  Node
	right Node
}

func (t Plus) eval(e Env) (int, error) {
	lvalue, err := t.left.eval(e)
	if err != nil {
		return 0.0, err
	}
	rvalue, err := t.right.eval(e)
	if err != nil {
		return 0.0, err
	}
	return lvalue + rvalue, nil
}

type Mult struct {
	left  Node
	right Node
}

func (t Mult) eval(e Env) (int, error) {
	lvalue, err := t.left.eval(e)
	if err != nil {
		return 0.0, err
	}
	rvalue, err := t.right.eval(e)
	if err != nil {
		return 0.0, err
	}
	return lvalue * rvalue, nil
}
