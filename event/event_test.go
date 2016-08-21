package event

//go:generate protoc --go_out=. event.proto

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

func TestEventThingStarted(t *testing.T) {
	tst := &EventThingStarted{
		Type: EventThingStarted_THING_ONE,
		Name: "Bob",
	}

	assert.NotNil(t, tst)
}

func TestEventThingStartedInEvent(t *testing.T) {
	tst := &EventThingStarted{
		Type: EventThingStarted_THING_ONE,
		Name: "Bob",
	}
	assert.NotNil(t, tst)

	a, err := ptypes.MarshalAny(tst)
	assert.NoError(t, err)
	assert.NotNil(t, a)

	e := &Event{
		Payload: a,
	}
	assert.NotNil(t, e)

	data, err := proto.Marshal(e)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	newe := &Event{}
	err = proto.Unmarshal(data, newe)
	assert.NoError(t, err)

	var x ptypes.DynamicAny
	err = ptypes.UnmarshalAny(newe.Payload, &x)
	assert.NoError(t, err)
	fmt.Printf("Got: %V\n", x)

	switch ev := x.Message.(type) {
	case *EventThingStarted:
		assert.Equal(t, "Bob", ev.Name)
	}
}
