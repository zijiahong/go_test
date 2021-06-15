package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// create and add some files to the archive
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"reade.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handing license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Panic(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Panic(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Panic(err)
	}

	// Open and iterate through the files in the archive
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // END of archive
		}
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Contents of %s \n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Panic(err)
		}
		fmt.Println()
	}
}
