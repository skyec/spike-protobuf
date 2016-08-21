package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate protoc --go_out=. event.proto

func TestEventThingStarted(t *testing.T) {
	tst := &EventThingStarted{
		Type: EventThingStarted_THING_ONE,
		Name: "Bob",
	}

	assert.NotNil(t, tst)
}
