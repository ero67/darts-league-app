package entities

// stringPtr returns a pointer to the string value
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// intPtr returns a pointer to the int value
func intPtr(i int) *int {
	return &i
}

// float64Ptr returns a pointer to the float64 value
func float64Ptr(f float64) *float64 {
	return &f
}