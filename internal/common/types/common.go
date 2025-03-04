package types

func createList[T any](validFields map[string]T, whitelist bool, targets ...string) map[string]T {
	list := make(map[string]T)

	if !whitelist {
		list = validFields
	}

	for _, t := range targets {

		if whitelist {
			if val, exists := validFields[t]; exists {
				list[t] = val
			}

		} else {
			delete(list, t)
		}
	}

	return list
}

func isExist[T any](target string, list map[string]T) bool {
	_, exists := list[target]
	return exists
}
