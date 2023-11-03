package utils

func GetValueOrEmptyString(val *string) string {
	if val != nil {
		return *val
	}
	return ""
}
