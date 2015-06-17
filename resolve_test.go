package resolve

import (
	"log"
	"os"
	"testing"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestResolve(t *testing.T) {
	pwd, err := os.Getwd()
	assert(err)

	_, err = os.Create(pwd + "/" + "hello.js")

	dependency, err := Resolve("./hello.js", pwd)
	if err != nil {
		t.Error("Should not have thrown an error: ", err)
	}
	actual := dependency.Pathname
	expected := pwd + string(os.PathSeparator) + "hello.js"
	if actual != expected {
		t.Error("Expected ", expected, " got ", actual)
	}

	assert(os.Remove(pwd + "/" + "hello.js"))
}
