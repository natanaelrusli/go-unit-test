package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Nata)",
			request: "Nata",
			expected: "Hello Nata",
		},
		{
			name: "HelloWorld(Nael)",
			request: "Nael",
			expected: "Hello Nael",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("befor unit test")

	m.Run()

	fmt.Println("after unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Nata", func(t *testing.T) {
		result := HelloWorld("Nata")
		require.Equal(t, "Hello Nata", result)
	})

	t.Run("Nael", func(t *testing.T) {
		result := HelloWorld("Nael")
		require.Equal(t, "Hello Nael", result)
	})
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("nael")
	assert.Equal(t, "Hello nael", result, "result must be 'Hello nael'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("nael")
	require.Equal(t, "Hello nael", result, "result must be 'Hello nael'")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa berjalan di mac")
	}

	result := HelloWorld("Nael")
	require.Equal(t, "Hello Nael", result)
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