package errors

type ErrorInvokingProgram struct{}

func (e ErrorInvokingProgram) Error() string {
	return "Error calling the program\nInvoke the program with -h for more information"
}
