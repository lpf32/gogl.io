package main

import "fmt"

func getError() ([]string, error) {
	return nil, fmt.Errorf("test error")
}

func main()  {
	_, err := getError()

	if err != nil {
		fmt.Println(err)
	}

}
