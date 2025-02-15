package main

import (
	"fmt"
	"strings"
)

// Convert relative path to absolute path
func ConvertToAbsolutePath(currentDir, relativePath string) string {
	stack := new(Stack)

	// Split paths into segments
	currentParts := strings.Split(currentDir, "/")
	relativeParts := strings.Split(relativePath, "/")

	// Process current directory
	for _, part := range currentParts {
		if part != "" {
			stack.Push(part)
		}
	}

	// Process relative path
	for _, part := range relativeParts {
		switch part {
		case "", ".":
			// Ignore empty and current directory (.)
		case "..":
			if !stack.IsEmpty() {
				stack.Pop()
			}
		default:
			stack.Push(part)
		}
	}

	// Construct absolute path
	absolutePath := "/" + strings.Join(stack.s, "/")
	return absolutePath
}

func main() {
	currentDir := "/home/user/docs/rouhan"
	relativePath := "./pictures/image.jpg"

	absolutePath := ConvertToAbsolutePath(currentDir, relativePath)
	fmt.Println("Absolute Path:", absolutePath)
}
