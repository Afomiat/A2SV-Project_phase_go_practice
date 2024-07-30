package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func checkPali(s string) bool{
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w\s]`)
	s = re.ReplaceAllString(s, "")

	new_s := []rune(s)
	n := len(new_s)

	for i := 0; i < n/2; i ++{
		new_s[i], new_s[n - i - 1] = new_s[n - i - 1], new_s[i]
	}

	if string(new_s) == s{
		return true
	}else{
		return false
	}

}

func main(){
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	val := checkPali(str)
	fmt.Println(val)
	if val{
		fmt.Printf("%v is a palindrome \n", str)
	}else{
		fmt.Printf("%v is not a palindrome \n", str)
	}

	
}