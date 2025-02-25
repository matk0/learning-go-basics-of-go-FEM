package main

import (
	"fmt"
	"os"
	"frontendmasters.com/go/files/fileutils"
)

func main() {
    rootPath, _ := os.Getwd()
    filePath := rootPath + "/data/text.txt"

    content, err := fileutils.ReadFile(filePath)

    if (err == nil) {
        fmt.Println(content)
        newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", content, content, content)
        fileutils.WriteToFile(filePath + ".output.txt", newContent)
    } else {
        fmt.Printf("ERROR: %s\n", err)
    }

}
