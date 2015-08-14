package gokatas

type OptionError struct {
	text string
}

type Fun func(Any) Any

func (this OptionError) Error() string { return this.text }

type Any interface{}

type Option interface {
	Get() (Any, OptionError)
	GetOrElse(Any) Any
	Map(Fun)
}

type None struct{}

func (this None) Get() (Any, error) {
	return nil, OptionError{text: "No value"}
}

func (this None) GetOrElse(e Any) Any {
	return e
}

func (this None) Map(f Fun) Any {
	return None{}
}

type Some struct {
	val Any
}

func (this Some) Get() (Any, error) {
	return this.val, nil
}

func (this Some) GetOrElse(e Any) Any {
	return this.val
}

func (this Some) Map(f Fun) Any {
	return Some{val: f(this.val)}
}
