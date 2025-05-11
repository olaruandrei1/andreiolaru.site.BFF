package helpers

import "strings"

func stringInSliceInsensitive(target string, list []string) bool {
	for _, s := range list {
		if strings.EqualFold(s, target) {
			return true
		}
	}
	return false
}
