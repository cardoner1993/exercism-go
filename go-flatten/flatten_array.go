package flatten

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{} // Empty slice, not nil

	switch v := nested.(type) {
	case []interface{}:
		for _, item := range v {
			if item != nil {
				if subList, ok := item.([]interface{}); ok {
					result = append(result, Flatten(subList)...)
				} else {
					result = append(result, item)
				}
			}
		}
	default:
		if nested != nil {
			result = append(result, nested)
		}
	}

	return result
}
