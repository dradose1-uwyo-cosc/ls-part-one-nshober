package functions

import (
	"os"
)

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	filter := []os.DirEntry{}
	for _, entry := range entries {
		if len(entry.Name()) > 0 && entry.Name()[0] != '.' {
			filter = append(filter, entry)
		}
	}
	return filter
}
