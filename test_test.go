package test

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
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
		Name: "yukpiz",
	}
	assert.Equal(t, u2, u3)
}

func Test_Example2(t *testing.T) {
	x1 := xerrors.New("aaaaa")

	log.Println(reflect.TypeOf(x1))
	assert.EqualError(t, x1, "aaaaa")
}
