package errors

type ErrorInvokingProgram struct{}

func (e ErrorInvokingProgram) Error() string {
	return "Error calling the program\nInvoke the program with -h for more information"
}

type ErrorTypeOfLengthPassword struct{}

func (e ErrorTypeOfLengthPassword) Error() string {
	return "You must pass an integer to the length password's arg"
}

type ErrorTypeOfRotations struct{}

func (e ErrorTypeOfRotations) Error() string {
	return "You must pass an integer to the rotations arg"
}

type ErrorHashNotAvailable struct{}

func (e ErrorHashNotAvailable) Error() string {
	return "Hash not available"
}
