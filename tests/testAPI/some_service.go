package testAPI

type SomeService struct {}

func NewRecursiveSerice(*SomeService) *SomeService {
	return &SomeService{}
}
