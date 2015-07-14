package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestServo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Servo Test Suite")
}
