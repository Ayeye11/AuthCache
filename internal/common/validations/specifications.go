package validations

type Specification struct {
	optional    bool
	minSize     int
	maxSize     int
	targetError error
	patterns    []string
}

func NewSpec(optional bool, minSize, maxSize int, targetError error, patterns ...string) *Specification {
	return &Specification{optional, minSize, maxSize, targetError, patterns}
}
