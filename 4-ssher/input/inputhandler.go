package inputhandler

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cetinboran/ssher/errorparser"
)

type IInputHandler interface {
	Start() ([]Arg, error)
}

type InputHandler struct {
	ErrorHandler *errorparser.ErrorHandler
	ValidOptions []Arg
}

func NewInputHandler() IInputHandler {
	var validOptions []Arg

	validOptions = append(validOptions, *NewArg("-u", "Known Username", true))
	validOptions = append(validOptions, *NewArg("-U", "Username Wordlist", true))
	validOptions = append(validOptions, *NewArg("-p", "Known Password", true))
	validOptions = append(validOptions, *NewArg("-P", "Password Wordlist", true))
	validOptions = append(validOptions, *NewArg("-s", "Server IP", true))
	validOptions = append(validOptions, *NewArg("--size", "Worker Count - Default 4", true))
	validOptions = append(validOptions, *NewArg("--port", "Port - Default 22", true))

	return &InputHandler{
		ValidOptions: validOptions,
		ErrorHandler: errorparser.NewErrorHandler(),
	}
}

func (i *InputHandler) Start() ([]Arg, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		i.Helper()
		os.Exit(0)
	}

	userOptions, err := i.FindOptions(args)
	if err != nil {
		return nil, err
	}

	if !i.ConflictChecker(userOptions) {
		return nil, i.ErrorHandler.Send(1, "Cannot use -p and -P together or -u and -U")
	}

	newUserOptions := i.AddDefaults(userOptions)

	if err := i.ErrorChecker(newUserOptions); err != nil {
		return nil, err
	}

	return newUserOptions, nil
}

func (i *InputHandler) FindOptions(args []string) (options []Arg, err error) {
	for j := 0; j < len(args); j++ {
		var newArg Arg
		if strings.Contains(args[j], "-") {
			if !i.ValidOption(args[j]) {
				return nil, i.ErrorHandler.Send(1, "Invalid option "+args[j])
			}
			if i.Has(options, args[j]) {
				return nil, i.ErrorHandler.Send(1, "This option "+args[j]+" already being used")
			}
			newArg.SetField(args[j])
		}

		if i.CheckNeedInput(args[j]) {
			j++
			if len(args) <= j {
				return nil, i.ErrorHandler.Send(2, "Need input for this option "+args[j-1])
			}
			newArg.SetInput(args[j])
			if strings.Contains(args[j], "-") {
				return nil, i.ErrorHandler.Send(2, "Need input for this option"+args[j])
			}
		}

		options = append(options, newArg)
	}
	return options, nil
}

func (i *InputHandler) ValidOption(arg string) bool {
	for _, validArg := range i.ValidOptions {
		if validArg.Field() == arg {
			return true
		}
	}

	return false
}

func (i *InputHandler) CheckNeedInput(arg string) bool {
	for _, validArg := range i.ValidOptions {
		if validArg.Field() == arg {
			return validArg.NeedInput()
		}
	}

	return false
}

func (i *InputHandler) Has(options []Arg, userArg string) bool {
	for _, option := range options {
		if option.Field() == userArg {
			return true
		}
	}

	return false
}

func (i *InputHandler) ConflictChecker(options []Arg) bool {
	var pUsed, PUsed bool
	var uUsed, UUsed bool

	for _, arg := range options {
		switch arg.Field() {
		case "-p":
			pUsed = true
		case "-P":
			PUsed = true
		case "-u":
			uUsed = true
		case "-U":
			UUsed = true
		}
	}

	if pUsed && PUsed {
		return false
	}
	if uUsed && UUsed {
		return false
	}
	return true
}

func (i *InputHandler) AddDefaults(options []Arg) []Arg {
	var newArg Arg
	if !i.Has(options, "--port") {
		newArg.SetField("--port")
		newArg.SetInput("22")

		options = append(options, newArg)
	}

	if !i.Has(options, "--size") {
		newArg.SetField("--size")
		newArg.SetInput("4")

		options = append(options, newArg)
	}

	return options
}

func (i *InputHandler) ErrorChecker(options []Arg) error {
	if !i.Has(options, "-s") {
		return i.ErrorHandler.Send(2, "Server IP Required")
	}

	if i.Has(options, "--size") {
		workerCountInt, err := strconv.Atoi(Get(options, "--size").Input())
		if err != nil {
			return err
		}

		if workerCountInt > 8 {
			return i.ErrorHandler.Send(1, "Worker count cannot be larger than 8")
		}
	}

	return nil
}

func (i *InputHandler) Helper() {
	fmt.Println("Usage of SSHER:")
	fmt.Println("  -u <username>      Specify a single username")
	fmt.Println("  -U <userlist>      Specify a file containing a list of usernames")
	fmt.Println("  -p <password>      Specify a single password")
	fmt.Println("  -P <passwordlist>  Specify a file containing a list of passwords")
	fmt.Println("\nExample:")
	fmt.Println("  SSHER -u admin -p 123456")
	fmt.Println("  SSHER -U users.txt -P passwords.txt")
}
