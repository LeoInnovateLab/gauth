package utils

// SafeStringConvert safely converts interface{} to string
// Returns empty string if value is nil or not a string type
func SafeStringConvert(value interface{}) string {
	if value == nil {
		return ""
	}
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

// SafeIntConvert safely converts interface{} to int
// Returns 0 if value is nil or cannot be converted to int
func SafeIntConvert(value interface{}) int {
	if value == nil {
		return 0
	}
	if num, ok := value.(float64); ok {
		return int(num)
	}
	if num, ok := value.(int); ok {
		return num
	}
	return 0
}
