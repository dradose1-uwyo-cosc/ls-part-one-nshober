package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func SimpleLS(w io.Writer, args []string, useColor bool) {
	directories := []string{}
	files := []string{}
	if len(args) == 0 {
		args = []string{"."}
	}
	for _, arg := range args {
		info, err := os.Lstat(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %s\n", arg, err.Error())
			continue
		}
		if info.IsDir() {
			directories = append(directories, arg)
		} else {
			files = append(files, arg)
		}

	}
	sort.Slice(files, func(i, j int) bool { return strings.ToLower(files[i]) < strings.ToLower(files[j])})
	for _, file := range files {
		info, _ := os.Lstat(file)
		mode := info.Mode()
		exec := mode.IsRegular() && (mode&0111) != 0
		if useColor && exec {
				Green.ColorPrint(w, filepath.Base(file)+"\n")
			} else {
				io.WriteString(w, filepath.Base(file)+"\n")
			}
	}

	if len(files) > 0 && len(directories) > 0 {
		io.WriteString(w, "\n")
	}

	current := ""
	others := []string{}
	for _, directory := range directories {
		if directory == "." {
			current = directory
		} else {
			others = append(others, directory)
		}
	}
	sort.Slice(others, func(i, j int) bool { return strings.ToLower((others[i])) < strings.ToLower(others[j])})
	allDirectories := []string{}
	if current != "" {
		allDirectories = append(allDirectories, current)
	}
	allDirectories = append(allDirectories, others...)
	for _, directory := range allDirectories {
		if len(allDirectories) > 1 && len(args) > 1 {
			if useColor {
				Blue.ColorPrint(w, filepath.Base(directory)+":\n")
			} else {
				io.WriteString(w, filepath.Base(directory)+":\n")
			}
		}
		entries, err := os.ReadDir(directory)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %s\n", directory, err.Error())
			continue
		}
		entries = dirFilter(entries)
		names := []string{}
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		sort.Slice(names, func(i, j int) bool { return strings.ToLower(names[i]) < strings.ToLower(names[j])})
		for _, name := range names {
			path := filepath.Join(directory, name)
			info, err := os.Lstat(path)
			mode := info.Mode()
			exec := mode.IsRegular() && (mode&0111) != 0
			if err != nil {
				fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %s\n", path, err.Error())
				continue
			}
			if useColor && exec {
				Green.ColorPrint(w, filepath.Base(name)+"\n")
			} else if useColor && info.IsDir() {
				Blue.ColorPrint(w, filepath.Base(name)+"\n")
			} else {
				io.WriteString(w, filepath.Base(name)+"\n")
			}
		}
		if len(allDirectories) > 1 && len(args) > 1 {
		io.WriteString(w, "\n")
		}
	}
}
