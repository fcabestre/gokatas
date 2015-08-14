package gokatas

import "testing"

func TestOption0(t *testing.T) {
	o := None{}
	_, e := o.Get()
	if e == nil {
		t.Fail()
	}
}

func TestOption1(t *testing.T) {
	o := Some{val: 5}
	v, _ := o.Get()
	if v != 5 {
		t.Fail()
	}
}

func TestOption2(t *testing.T) {
	o := None{}
	v := o.GetOrElse(8)
	if v != 8 {
		t.Fail()
	}
}

func TestOption3(t *testing.T) {
	o := Some{val: 5}
	v := o.GetOrElse(8)
	if v != 5 {
		t.Fail()
	}
}

func TestOption4(t *testing.T) {
	var f Fun = func(a Any) Any { return a.(int) + 4 }
	o := None{}
	v := o.Map(f)
	expected := None{}
	if v != expected {
		t.Fail()
	}
}

func TestOption5(t *testing.T) {
	var f Fun = func(a Any) Any { return a.(int) + 4 }
	o := Some{val: 5}
	v := o.Map(f)
	expected := Some{val: 9}
	if v != expected {
		t.Fail()
	}
}
