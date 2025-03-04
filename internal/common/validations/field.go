package validations

import "github.com/Ayeye11/se-thr/internal/common/errs"

func ValidateField(target any, spec *Specification) error {
	if spec == nil {
		return errs.InternalX(errs.BscError("missing field specifications"))
	}

	if spec.optional && target == nil {
		return nil
	}

	if target == nil {
		return errs.MissingCredentials
	}

	switch t := target.(type) {

	case string:
		if spec.optional && t == "" {
			return nil
		}

		if !validateLength(t, spec.minSize, spec.maxSize) {
			return spec.targetError
		}

		for _, p := range spec.patterns {
			if !validatePattern(t, p) {
				return spec.targetError
			}
		}

		return nil

	case *string:
		if spec.optional && t == nil {
			return nil
		}

		if !validateLength(t, spec.minSize, spec.maxSize) {
			return spec.targetError
		}

		for _, p := range spec.patterns {
			if !validatePattern(*t, p) {
				return spec.targetError
			}
		}

		return nil

	case int:
		if !validateLength(t, spec.minSize, spec.maxSize) {
			return spec.targetError
		}

		return nil

	case *int:
		if spec.optional && t == nil {
			return nil
		}

		if !validateLength(t, spec.minSize, spec.maxSize) {
			return spec.targetError
		}

		return nil

	default:
		return spec.targetError
	}
}
