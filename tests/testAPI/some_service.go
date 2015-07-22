package testAPI

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
