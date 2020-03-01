package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// main function
// need arguments source destination
func main() {
	if len(os.Args) != 3 {
		fmt.Println("App need source and destination")

		// filepath.Base(os.Args[0])) get the current path
		fmt.Printf("usage: %s source destination\n", filepath.Base(os.Args[0]))
		return
	}

	// get source and check if file exist
	source := os.Args[1]
	fmt.Printf("The source file is %s\n", source)

	//check if file exist
	sourceExist, fileSize := fileExists(source)
	if sourceExist {
		fmt.Printf("File dose exist = %t\nFilesize is %d\n", sourceExist, fileSize)
	}

	// get the destination and create new file if not exist
	destination := os.Args[2]
	fmt.Printf("The destination file is %s\n", destination)

	// try to coyp the file
	errorCopy := copyFile(source, destination)
	if os.IsNotExist(errorCopy) {
		fmt.Printf("copy file failed: %q\n", errorCopy)
	}
}

// CopyFile return an error if the copy function fail
// else the error is nil
func copyFile(src, dst string) error {
	return nil
}

/*
Source from Opensource.com -- see Readme.md for more information
func copy(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
*/

// fileIsFolder check if the file exist and if the file is a folder
// return bool
func fileIsFolder(filename string) bool {
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	info, error := os.Stat(filename)

	// check if the file exist
	if os.IsNotExist(error) {
		return false
	}

	return info.IsDir()
}

// fileExists return bool and int64
// bool If the file exist and is not a folder
// int64 Size of the file -> file size is 0, if the file don't exist
func fileExists(filename string) (bool, int64) {
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	info, error := os.Stat(filename)

	// check if the file exist
	if os.IsNotExist(error) {
		return false, 0
	}

	// check if the file object is a folder
	return !info.IsDir(), info.Size()
}
