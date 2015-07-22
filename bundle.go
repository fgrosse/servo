package servo

type Bundle interface {
	Boot(*Kernel)
}
