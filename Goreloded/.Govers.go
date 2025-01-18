package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//////////////////////( Open file )//////////////////////////
	var fileName string
	fmt.Print("Enter the name of the file: ")
	fmt.Scanln(&fileName)
	file, err := ft_openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//////////////////////( Reading from file )//////////////////////////
	var lines []string
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if len(line) > 0 {
				lines = append(lines, line)
			}
			break
		}
		lines = append(lines, line)
	}
	line := strings.Join(lines, "")
	
	//////////////////////( writing to file )//////////////////////////
	file1, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Opps Error in Creating", err)
		return
	}
	defer file1.Close()
	data := []byte(line)
	_, err = file1.Write(data)
	if err != nil {
		fmt.Println("Opps Error in Writing", err)
		return
	}
	fmt.Println("Data Written Successfully")
}

//////////////////(Helpers Functions)//////////////////////////////
func ft_openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}


// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	//////////////////////( Reading from file )//////////////////////////
// 	// file_name := "file.txt"
	
// 	var fileName string
// 	fmt.Print("Enter the name of the file: ")
// 	fmt.Scanln(&fileName)
// 	file, err := ft_openFile(fileName)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	rd := bufio.NewReader(file)
// 	var line string
// 	for {
// 		line, err = rd.ReadString('\n')
// 		fmt.Println(line)
// 		if err != nil {
// 			break
// 		}
// 	}
// 	// fmt.Println("Buuuuugg")

// 	//////////////////////( writing to file )//////////////////////////
// 	file1 , err := os.Create("output.txt")
// 	if err != nil {
// 		fmt.Println("Opps Error in Creating", err)
// 		return
// 	}
// 	defer file1.Close()
// 	data := []byte(line)
// 	_, err = file1.Write(data)
// 	if err != nil {
// 		fmt.Println("Opps Error in Writing", err)
// 		return
// 	}
// 	fmt.Println("Data Written Successfully")
// }

// func ft_openFile(fileName string) (*os.File, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening file: %v", err)
// 	}
// 	return file, nil
// }