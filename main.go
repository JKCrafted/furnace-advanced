package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/3096/furnace/commands"
)

func multiItems() {
	dir, err := os.Open("Input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".") == false {
			erra := commands.ReplaceTexturesInWismt("Input/"+file.Name()+"/"+file.Name()+".wismt", "Replace/"+file.Name(), "Output/"+file.Name()+"/"+file.Name()+".wismt")
			if erra != nil {
				os.MkdirAll("Output/"+file.Name(), os.ModePerm)
				erra := commands.ReplaceTexturesInWismt("Input/"+file.Name()+"/"+file.Name()+".wismt", "Replace/"+file.Name(), "Output/"+file.Name()+"/"+file.Name()+".wismt")
				if erra != nil {
					fmt.Println(erra)
				}
			}
		}
	}
}

func multiVersion() {
	runs, err := strconv.Atoi(os.Args[2])
	fmt.Println(err, runs)
	if err == nil {
		for i := 1; i <= runs; i++ {
			dir, err := os.Open("Input")
			if err != nil {
				fmt.Println("Error:", err)
			}
			defer dir.Close()
			files, err := dir.Readdir(-1)
			if err != nil {
				fmt.Println("Error:", err)
			}
			erra := commands.ReplaceTexturesInWismt("Input/"+files[1].Name(), "Replace/"+strconv.Itoa(i), "Output/"+strconv.Itoa(i)+"/"+files[1].Name())
			if erra != nil {
				os.MkdirAll("Output/"+strconv.Itoa(i), os.ModePerm)
				erra := commands.ReplaceTexturesInWismt("Input/"+files[1].Name(), "Replace/"+strconv.Itoa(i), "Output/"+strconv.Itoa(i)+"/"+files[1].Name())
				if erra != nil {
					fmt.Println(erra)
				}
			}
		}
	}
}

func multiVerItems() {
	runs, err := strconv.Atoi(os.Args[2])
	if err == nil {
		for i := 1; i <= runs; i++ {
			dir, err := os.Open("Input")
			if err != nil {
				fmt.Println("Error:", err)
			}
			defer dir.Close()
			files, err := dir.Readdir(-1)
			if err != nil {
				fmt.Println("Error:", err)
			}
			for _, file := range files {
				if strings.Contains(file.Name(), ".") == false {
					erra := commands.ReplaceTexturesInWismt("Input/"+file.Name()+"/"+file.Name()+".wismt", "Replace/"+strconv.Itoa(i)+"/"+file.Name(), "Output/"+strconv.Itoa(i)+"/"+file.Name()+"/"+file.Name()+".wismt")
					if erra != nil {
						os.MkdirAll("Output/"+strconv.Itoa(i)+"/"+file.Name(), os.ModePerm)
						erra := commands.ReplaceTexturesInWismt("Input/"+file.Name()+"/"+file.Name()+".wismt", "Replace/"+strconv.Itoa(i)+"/"+file.Name(), "Output/"+strconv.Itoa(i)+"/"+file.Name()+"/"+file.Name()+".wismt")
						if erra != nil {
							fmt.Println(erra)
						}
					}
				}
			}

		}
	}
}

func singleItem() {
	dir, err := os.Open("Input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	erra := commands.ReplaceTexturesInWismt("Input/"+files[1].Name(), "Replace", "Output/"+files[1].Name())
	if erra != nil {
		os.MkdirAll("Output", os.ModePerm)
		erra := commands.ReplaceTexturesInWismt("Input/"+files[1].Name(), "Replace", "Output/"+files[1].Name())
		if err != nil {
			fmt.Println(erra)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		singleItem()
	} else {
		if os.Args[1] == "mi" {
			multiItems()
		} else if os.Args[1] == "mv" {
			multiVersion()
		} else if os.Args[1] == "mimv" || os.Args[1] == "mvmi" {
			multiVerItems()
		} else {
			singleItem()
		}
	}
	os.Exit(1)
}
