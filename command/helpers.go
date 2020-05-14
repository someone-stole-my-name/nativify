package command

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
