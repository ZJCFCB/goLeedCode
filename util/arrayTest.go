package util

type RunTest interface {
	Run()
}

func ArrayTest(run RunTest) {
	run.Run()
}
