package address

func GetPointer(i int) int {
	i = 50

	return i
}

func SetValue(i *int) {
	*i = 50
}
