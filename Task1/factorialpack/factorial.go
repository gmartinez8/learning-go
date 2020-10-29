package factorialpack

//Factorial First letter in function  should be uppercase for function to be exported
func Factorial(number uint64) uint64 {
	if number == 0 {
		return 1
	}
	return number * Factorial(number-1)
}
