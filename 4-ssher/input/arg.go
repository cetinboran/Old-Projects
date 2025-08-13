package inputhandler

type Arg struct {
	field     string
	input     string
	needInput bool
}

func NewArg(
	field, input string,
	needInput bool,
) *Arg {
	var newArg Arg

	newArg.SetField(field)
	newArg.SetInput(input)
	newArg.SetNeedInput(needInput)

	return &newArg
}

func Get(args []Arg, fied string) *Arg {
	for _, arg := range args {
		if arg.Field() == fied {
			return &arg
		}
	}

	return nil
}

// GETTER
func (a *Arg) Field() string {
	return a.field
}

func (a *Arg) NeedInput() bool {
	return a.needInput
}

func (a *Arg) Input() string {
	return a.input
}

// SETTER
func (a *Arg) SetField(fieldRef string) {
	a.field = fieldRef
}

func (a *Arg) SetInput(inputRef string) {
	a.input = inputRef
}

func (a *Arg) SetNeedInput(needInputRef bool) {
	a.needInput = needInputRef
}
