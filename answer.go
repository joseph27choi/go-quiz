package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	quizName := "problems.csv"

	// Open CSV
	// declares two variables: file and err
	// file represents the pointer that will point to the csv file in hardware memory
	// err represents any errors that might occur while opening the file
	// if it catches an error, then err will have a value indicating an error (maybe a return 1??)

	file, err := os.Open(quizName)

	// if the err variable has an error value, then it needs to notify the user that there was an error
	if err != nil {
		// it is the return value!!!
		fmt.Println("Error: ", err)
		return
	}
	// before moving on, I need to close the file that I just "attempted" to open 
	// when do I do this? after the function is completed
	// that is why this is called defer, just like a <script>
	defer file.Close()

	// I need to create a CSV reader
	// but why??
	// just like in c, must open then read
	// this step requires more code than python because I have to open then read/write
	// in python, you take care of everything with the keyword 'with'
	reader := csv.NewReader(file)

	// to read the file, I must do one extra step for some reason
	// go through row by row
	// apparently not deprecated?
	rows, err := reader.ReadAll()
	if (err != nil) {
		fmt.Print("Error: ", err)
		return
	}

	// in Go, the for loops are enumerated by default??? (index, elementValue) ... (array)
	// this is for maps
	// it also supports the c way of for loops
	// use underscore for the first parameter because don't need the index
	var i string

	for _, row := range rows {
		equation := row[0]
		corr_ans := row[1]

		// query user for their answer
		fmt.Printf("What is %s ? ", equation)
		fmt.Scan(&i)

		if i == corr_ans {
			fmt.Println("You are correct!")
		} else {
			fmt.Println("You are incorrect!")
			fmt.Printf("The equation %s has the answer %s\n", equation, corr_ans)
		}
	}



	// now that the corresponding answer


}




