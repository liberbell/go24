package main

import testing

func TestRun(t *tesing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
