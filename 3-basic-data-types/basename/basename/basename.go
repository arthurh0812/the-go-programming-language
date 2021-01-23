package basename

import "strings"

// Basename1 removes directory components and a .suffix.
func Basename1(s string) string {
	// discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

// Basename2 removes directory components and a .suffix.
func Basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
