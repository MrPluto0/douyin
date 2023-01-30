package check

func CheckPanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
