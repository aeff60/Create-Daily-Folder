package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	now := time.Now()
	formattedDate := now.Format("02_01_2006")

	folderName := fmt.Sprintf(formattedDate)

	if err := os.Mkdir(folderName, 0755); err != nil {
		fmt.Println("Error creating folder:", err)
	} else {
		fmt.Println("Folder created successfully:", folderName)
	}
}
