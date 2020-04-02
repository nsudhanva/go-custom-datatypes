package organisation

type Identifiable interface {
	ID() string
}

// No need to specify datatype of the interface
type Person struct {
}

// Method has the same signature as identifiable

func (p Person) ID() string {
	return "12345"
}
