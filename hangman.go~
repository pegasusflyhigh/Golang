package main
 
 import (
 	"fmt"
 	valid "github.com/asaskevich/govalidator"
 	"strings"
 )
 
 func hangman(str string) {
 
 	guessed := []string{}
 	guess := []string{}
 
 	for _ = range str {
 		guess = append(guess, "_")
 	}
 
 	var r string
 	lives := 6
 	for lives >= 1 {
 
 		fmt.Printf("\nyour lives : %d\n", lives)
 		fmt.Println(guess)
 		ctr, ctr1 := 0, 0
 		fmt.Printf("Enter your guess : ")
 		fmt.Scanf("%s", &r)
 		if !valid.IsAlpha(r) || len(r) > 1 {
 			fmt.Println("Please enter single alphabets!")
 		} else {
 			r = strings.ToLower(r)
 
 			if isinclude(guessed, r) {
 				fmt.Println("ALREADY GUESSED!")
 			} else {
 				guessed = append(guessed, r)
 
 				for i, _ := range str {
 					if string(str[i]) == (r) {
 						guess[i] = r
 						ctr = ctr + 1
 					}
 
 				}
 
 				if isinclude(guess, "_") {
 					ctr1 = ctr1 + 1
 				} else {
 
 					fmt.Println("\nYOU WON!!\n")
 					fmt.Printf("The word is  : %s\n", str)
 					lives = 0
 				}
 
 				if ctr == 0 {
 					lives = lives  1
 				}
 
 			}
 
 		}
 
 	}
 }
 
 func isinclude(word []string, v string) bool {
 	ctr := 0
 	for i, _ := range word {
 		if word[i] == v {
 			ctr = ctr + 1
 		}
 	}
 
 	if ctr == 0 {
 		return false
 	}
 	return true
 }
 
 func main() {
 
 	c := string("Google")
 	c = strings.ToLower(c)
 	//display(c)
 	hangman(c)

 }
