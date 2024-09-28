package cmd

import (
	"fmt"
	"os"
	"sync"
)

type Command struct {
	Arguments []string
}

func (c *Command) Pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}

	fmt.Println(dir)
}

func (c *Command) Ls() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}

	ls, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error while reading current directory")
		return
	}

	for _, v := range ls {
		fmt.Print(v.Name(), " ")
	}
	fmt.Print("\n")
}

func (c *Command) Touch() {
	if len(c.Arguments) == 0 {
		fmt.Println("Missing file name")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(c.Arguments))

	fileNames := c.Arguments
	for _, v := range fileNames {
		go func(v string) {
			defer wg.Done()

			file, err := os.Create(v)
			if err != nil {
				fmt.Printf("Error creating file %v\n", file)
				return
			}
			fmt.Printf("File %v created\n", file.Name())

			file.Close()
		}(v)
	}

	wg.Wait()
}

func (c *Command) Mkdir() {
	if len(c.Arguments) == 0 {
		fmt.Println("Missing directory name")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(c.Arguments))

	dirNames := c.Arguments
	for _, v := range dirNames {
		go func(v string) {
			err := os.Mkdir(v, 700)
			if err != nil {
				exists := os.IsExist(err)
				if exists {
					fmt.Printf("Directory %v already exists\n", v)
				} else {
					fmt.Printf("Error while creating directory %v", v)
				}
			}

			wg.Done()
		}(v)
	}

	wg.Wait()
}
