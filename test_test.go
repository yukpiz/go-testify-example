package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID   int
	Name string
}

func Test_Example1(t *testing.T) {
	u1 := &User{
		ID:   1,
		Name: "yukpiz",
	}

	u2 := &User{
		ID:   1,
		Name: "yukpiz",
	}

	assert.Equal(t, u1, u2)

	u3 := &User{
		ID:   1,
		Name: "yukpiz2",
	}
	assert.Equal(t, u2, u3)
}
