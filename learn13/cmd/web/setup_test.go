package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type myhandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
