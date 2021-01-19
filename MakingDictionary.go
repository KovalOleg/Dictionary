package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Word struct{
	line string
	files string
	counter int
}

func main() {

	var dictionary []string //Main dictionary
	var text string

	// Declare files
	file1, err := os.Open("Going Postal.txt")
	file2, err := os.Open("Interesting Times.txt")
	file3, err := os.Open("Making Money.txt")
	file4, err := os.Open("Monstrous Regiment.txt")
	file5, err := os.Open("Moving Pictures.txt")
	file6, err := os.Open("Sourcery.txt")
	file7, err := os.Open("The Color of Magic.txt")
	file8, err := os.Open("The Last Continent.txt")
	file9, err := os.Open("The Light Fantastic.txt")
	file10, err := os.Open("The Truth.txt")


	// Reading data from files
	text = readFromFile(file1, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file2, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file3, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file4, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file5, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file6, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file7, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file8, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file9, err)
	getArray(text, &dictionary)
	text = ""
	text = readFromFile(file10, err)
	getArray(text, &dictionary)
	text = ""

	// Creating exit files
	file, err := os.Create("Dictionary.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writeIn(file, &dictionary)

	fmt.Println("Count of words in dictionary: ",len(dictionary))
	stat, err := file.Stat()
	fmt.Println("Size of output file is",stat.Size()/1024,"KB")
}


//Writing info in dictionary files
func writeIn(file *os.File, dictionary *[]string){
	for _, value := range *dictionary{
		var text = ""
		text+=value+"\n"
		file.WriteString(text)
	}
}

//Deleting similar words
func clearAr(array *[]string){
	sort.Strings(*array)
	var ar []string

	var arLen = 0
	for i, value := range *array{
		if i == 0{
			ar = append(ar, value)
			arLen++
		} else if string(ar[arLen-1]) != value{
			ar = append(ar, value)
			arLen++
		}
	}
	*array = ar
}

// Reading data from file and putting it in
func readFromFile(file *os.File, err error) string{
	var text = ""
	if err != nil {
		fmt.Println(err)
		return "Wrong File!"
	}

	defer file.Close()
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		text += string(data[:n])
	}
	text += " "
	return text
}

func getArray(text string, dictionary *[]string) {
	var i = 0
	for {
		var word = ""
		if (text[i] >= 65 && text[i] <= 90) || (text[i] >= 97 && text[i] <= 122) {
			word, i = getWord(i, text)
			*dictionary = append(*dictionary, strings.ToLower(word))
		} else {
			i++
		}
		if i >= len(text) {
			break
		}
	}
	clearAr(dictionary)



}

func getWord(i int, text string) (string, int) {
	var word = ""
	for {
		if (text[i] >= 65 && text[i] <= 90) || (text[i] >= 97 && text[i] <= 122) {
			word += string(text[i])
			i++
		} else if text[i] == 226 {
			word += ""
			i++
		} else if text[i] == 45 {
			word += "-"
			i++
		} else {
			break
		}
	}
	return word, i
}
