package stacknqueue

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var stack []string

	for _, p := range paths {
		switch p {
		case "", ".":
			continue
		case "..":
			if len(stack) >= 1 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, p)
		}
	}

	return "/" + strings.Join(stack, "/")
}
