package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*
"Getting file information"

"This program will print the information about a file, namely when it was
last modified, who owns it, how many bytes it is, and what its permissions are."
*/

func GetFileInfo() {
	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		// Print file info if it's a file (not a directory)
		if !info.IsDir() {
			fmt.Println("File name:", info.Name())
			fmt.Println("Size in bytes:", info.Size())
			fmt.Println("Permissions:", info.Mode())
			fmt.Println("Last modified:", info.ModTime())
			fmt.Println("Is Directory: ", info.IsDir())
			fmt.Printf("System interface type: %T\n", info.Sys())
			fmt.Printf("System info: %+v\n\n", info.Sys())
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}