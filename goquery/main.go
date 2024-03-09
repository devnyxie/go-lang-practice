package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the JSX file
	// file, err := os.Open("example.jsx")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// Parse the JSX file with goquery
	// doc, err := goquery.NewDocumentFromReader(file)
	// if err != nil {
	// 	fmt.Println("Error parsing file:", err)
	// 	return
	// }

	// For example, you can extract text content from JSX elements
	// doc.Find("div").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println("Text content of div:", strings.TrimSpace(s.Text()))
	// })

	// --- Add <p> tag to 1st div
	// content, err := os.ReadFile("example.jsx")
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// modifiedContent := string(content)
	// modifiedContent = strings.Replace(modifiedContent, "</div>", "<p>Additional text added</p></div>", -1)
	// err = os.WriteFile("example.jsx", []byte(modifiedContent), 0644) // 0644 is a file mode commonly used for permissions when creating or modifying files.
	// if err != nil {
	// 	fmt.Println("Error writing modified content to file:", err)
	// 	return
	// }

	// Read the content of the JSX file
	content, err := os.ReadFile("example.jsx")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert content to string
	modifiedContent := string(content)

	// Define the HTML content to insert between the comments
	htmlToInsert := "<p>This is the HTML content to insert</p>"

	// Find the position of the start and end comments
	startIndex := strings.Index(modifiedContent, "{/* CONTENT START */}")
	endIndex := strings.Index(modifiedContent, "{/* CONTENT END */}")

	// If both start and end comments are found
	if startIndex != -1 && endIndex != -1 {
		// Insert the HTML content between the start and end comments
		modifiedContent = modifiedContent[:startIndex+len("{/* CONTENT START */}")] +
			htmlToInsert +
			modifiedContent[endIndex:]
	} else {
		fmt.Println("Start or end comment not found")
		return
	}

	// Write the modified content back to the file
	err = os.WriteFile("example.jsx", []byte(modifiedContent), 0644)
	if err != nil {
		fmt.Println("Error writing modified content to file:", err)
		return
	}

	fmt.Println("Modified content saved to example.jsx")
}
