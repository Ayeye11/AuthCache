package validations

import "regexp"

func validatePattern(target, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(target)
}

func validateLength(target any, min, max int) bool {
	if target == nil {
		return false
	}

	switch t := target.(type) {

	case string:
		size := len(t)
		return min <= size && size <= max

	case *string:
		if t == nil {
			return false
		}

		size := len(*t)
		return min <= size && size <= max

	case int:
		return min <= t && t <= max

	case *int:
		if t == nil {
			return false
		}

		return min <= *t && *t <= max

	default:
		return false
	}
}
