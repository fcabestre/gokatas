package gokatas
import "testing"

func TestCalc0(t *testing.T) {
	e := Env{"x": 3}
	n := Litt{value:1}
	if v,_ := n.eval(e); v != 1 { t.Fail() }
}

func TestCalc1(t *testing.T) {
	e := Env{"x": 3}
	n := Var{name:"x"}
	if v,_ := n.eval(e); v != 3 { t.Fail() }
}

func TestCalc2(t *testing.T) {
	e := Env{"x": 3}
	n := Var{name:"y"}
	if _,e := n.eval(e); e == nil { t.Fail() }
}

func TestCalc3(t *testing.T) {
	e := Env{"x": 3}
	n := Plus{left:Var{name:"x"}, right:Litt{value:1}}
	if v,_ := n.eval(e); v != 4 { t.Fail() }
}

func TestCalc4(t *testing.T) {
	e := Env{"x": 3}
	n := Plus{right:Litt{value:1}, left:Var{name:"x"}}
	if v,_ := n.eval(e); v != 4 { t.Fail() }
}

func TestCalc5(t *testing.T) {
	e := Env{"x": 3}
	n := Mult{left:Var{name:"x"}, right:Litt{value:1}}
	if v,_ := n.eval(e); v != 3 { t.Fail() }
}

func TestCalc6(t *testing.T) {
	e := Env{"x": 3}
	n := Mult{right:Litt{value:1}, left:Var{name:"x"}}
	if v,_ := n.eval(e); v != 3 { t.Fail() }
}

func TestCalc7(t *testing.T) {
	e := Env{"x": 3, "y": 7}
	n := Plus{left:Mult{right:Litt{value:1}, left:Var{name:"x"}},right:Var{name:"y"}}
	if v,_ := n.eval(e); v != 10 { t.Fail() }
}

