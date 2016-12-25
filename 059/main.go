package main

import (
	"fmt"
	"io/ioutil"
	_ "strconv"
	"strings"
)

func asciiToString(encoded string, key [3]int) string {
	decoded := ""
	encoded = strings.Trim(encoded, " \n")

	for i, v := range strings.Split(encoded, ",") {
		var charVal int

		keyVal := key[i%3]
		fmt.Sscanf(v, "%d", &charVal)

		decodedVal := charVal ^ keyVal
		decoded += string(decodedVal)
	}

	return decoded
}

func main() {
	fileContentsBytes, _ := ioutil.ReadFile("p059_cipher.txt")
	encodedMessage := string(fileContentsBytes)

	/*
		for i := 97; i <= 122; i++ {
			for j := 97; i <= 122; j++ {
				for k := 97; i <= 122; k++ {
					// key := string(i) + string(j) + string(k)
					key := [3]int{i, j, k}
					decodedMessage := asciiToString(encodedMessage, key)

					if strings.Contains(decodedMessage, "the") && strings.Contains(decodedMessage, "an") {
						fmt.Println("key = ", key)
						fmt.Println("message = ", decodedMessage)
					}
				}
			}
		}
	*/

	sum := 0
	decodedMessage := asciiToString(encodedMessage, [3]int{103, 111, 100})

	for _, v := range decodedMessage {
		sum += int(v)
	}

	fmt.Println(decodedMessage)
	fmt.Println(sum)
}
