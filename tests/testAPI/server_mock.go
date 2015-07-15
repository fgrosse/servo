package testAPI

import (
	"fmt"

	. "github.com/onsi/gomega"
)

// ServerMock is a non blocking noop implementation of servo.Server
type ServerMock struct {
	RunHasBeenCalled bool
	ReturnError      bool

	Parameter1, Parameter2 string
}

func NewServerMockWithParams(param1, param2 string) *ServerMock {
	Expect(param1).To(Equal("foo"), `NewServerMockWithParams should always be called with the values "foo" and "bar"`)
	Expect(param2).To(Equal("bar"), `NewServerMockWithParams should always be called with the values "foo" and "bar"`)

	return &ServerMock{
		Parameter1: param1,
		Parameter2: param2,
	}
}

func (s *ServerMock) Run() error {
	s.RunHasBeenCalled = true
	if s.ReturnError {
		return fmt.Errorf("ServerMock was told to return an error!")
	}

	return nil
}
