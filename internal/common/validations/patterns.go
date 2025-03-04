package validations

const (
	PatternEmail      = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	PatternBcryptHash = `^\$2[ayb]\$\d{2}\$[A-Za-z0-9./]{53}$`
)
