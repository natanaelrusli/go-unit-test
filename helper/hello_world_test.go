package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("nael")
	assert.Equal(t, "Hello nael", result, "result must be 'Hello nael'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nael")
	if result != "Hello Nael" {
		t.Error("result must be Hello Nael")
	}
	fmt.Println("done executing TestHelloWorld")
}

func TestHelloWorldNata(t *testing.T) {
	result := HelloWorld("Nata")
	if result != "Hello Nata" {
		t.Fatal("result must be Hello Nata")
	}
	fmt.Println("done executing TestHelloWorldNata")
}