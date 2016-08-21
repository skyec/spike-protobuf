package addressbook

//go:generate protoc --go_out=. addressbook.proto

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestAddressBook(t *testing.T) {
	ab := AddressBook{}
	assert.NotNil(t, ab)
}

func TestPerson(t *testing.T) {
	p := Person{}
	assert.NotNil(t, p)
}

func TestWriteAddressBook(t *testing.T) {
	p := Person{
		Id:    123,
		Name:  "Me",
		Email: "me@me.com",
		Phones: []*Person_PhoneNumber{
			{Number: "000-000-0000", Type: Person_WORK},
		},
	}
	assert.NotNil(t, p)

	ab := AddressBook{
		People: []*Person{&p},
	}
	assert.NotNil(t, ab)

	out, err := proto.Marshal(&ab)
	assert.NoError(t, err)
	assert.NotNil(t, out)

	newbook := &AddressBook{}
	err = proto.Unmarshal(out, newbook)
	assert.NoError(t, err)
	assert.NotNil(t, newbook.People)
	assert.Len(t, newbook.People, 1)
	assert.Equal(t, "Me", newbook.People[0].Name)
}
