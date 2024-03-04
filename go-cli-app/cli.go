package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Structs
type Component struct {
	name string
}

type Page struct {
	name       string
	components []Component
}

type Theme struct {
	name  string
	pages []Page //default pages
}

type Config struct {
	theme Theme
}

// Default Themes
var themes = []Theme{
	{
		name: "Nigiri",
		pages: []Page{
			{
				name: "Home",
				components: []Component{
					{
						name: "profile",
					},
					{
						name: "blog",
					},
				},
			},
			{
				name: "About Me",
				components: []Component{
					{
						name: "about_me",
					},
				},
			},
		},
	},
	{
		name: "Minimal",
		pages: []Page{
			{
				name: "Home",
				components: []Component{
					{
						name: "profile",
					},
					{
						name: "portfolio",
					},
					{
						name: "blog",
					},
				},
			},
		},
	},
}

// Functions
func askTheme() (Theme, error) {
	var themeNames []string

	for _, theme := range themes {
		themeNames = append(themeNames, theme.name)
	}

	prompt := promptui.Select{
		Label: "Select starting template",
		Items: themeNames,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return Theme{}, err
	}
	// find theme by theme.name in themes slice (result is a theme name)
	for _, theme := range themes {
		if theme.name == result {
			return theme, nil
		}
	}
	fmt.Println("Theme not found")
	return Theme{}, err
}

func askPages(theme *Theme) error {
	var pagesNames []string

	for _, page := range theme.pages {
		fmt.Println(page.name)
		pagesNames = append(pagesNames, page.name)
	}
	//add "add page" functionality
	pagesNames = append(pagesNames, "Add Page")
	pagesNames = append(pagesNames, "Continue")
	prompt := promptui.Select{
		Label: "Select pages you would like to modify, continue when ready",
		Items: pagesNames,
	}

	_, result, err := prompt.Run()
	// to-do: add modify-page functionality
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	if result == "Add Page" {
		//add page
		fmt.Println("Add Page")
		prompt := promptui.Prompt{
			Label: "Enter new page's name",
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		theme.pages = append(theme.pages, Page{name: result})
		askPages(theme)
	} else if result == "Continue" {
		return nil
	} else {
		fmt.Println("Modify Page")
		//modify page
		// 1. Del page
		// 2. Edit page's name
		// 3. Del/Add page's components
		// so we have 3 buttons.
	}
	//to-do
	//0. "add page" functionality [done]
	//1. add "Continue" btn [done]
	//2. add modify-page functionality (add/delete components)
	//3. add "return" btn
	return nil
}

func askComponents(page Page) error {
	// var componentsNames []string

	// for _, component := range page.components {
	// 	componentsNames = append(componentsNames, component.name)
	// }
	// //add "add component" functionality
	// componentsNames = append(componentsNames, "Abort")
	// prompt := promptui.Select{
	// 	Label: "Select components you would like to modify",
	// 	Items: componentsNames,
	// }

	// _, result, err := prompt.Run()
	// // to-do: add modify-page functionality
	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return err
	// }

	return nil
}

// --- Main ---
func main() {
	//create empty config
	var config Config
	//ask for theme
	themeName, themeNameErr := askTheme()
	if themeNameErr != nil {
		return
	}
	config.theme = themeName
	// ask for pages
	askPagesErr := askPages(&config.theme)
	if askPagesErr != nil {
		return
	}
	// --- summary ---
	fmt.Printf("\nChosen Theme: %s\n", config.theme.name)
	fmt.Println("Chosen Pages:")
	for i, page := range config.theme.pages {
		fmt.Printf("%d. %s\n", i, page.name)
		fmt.Println("  Components:")
		for i, component := range page.components {
			fmt.Printf("  %d. %s \n", i, component.name)
		}
	}

}
