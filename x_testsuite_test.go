package servo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/fgrosse/servo"
	"fmt"
)

func TestServo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Servo Test Suite")
}

type TestBundle struct {}

func (b *TestBundle) Boot(kernel *servo.Kernel) {
	kernel.RegisterType("test_bundle.my_type", NewService)
}

type SomeService struct {}

func NewRecursiveService(*SomeService) *SomeService {
	return &SomeService{}
}

func NewService() *SomeService {
	return &SomeService{}
}

func NewServiceWithParam(param interface{}) *SomeService {
	panic(param)
	return &SomeService{}
}

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
