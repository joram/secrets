// Code generated by goderive DO NOT EDIT.

package main

// deriveContains returns whether the item is contained in the list.
func deriveContains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
