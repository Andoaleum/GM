package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
)

func Hello(name string) (string,error){
	if(name == ""){
		return "", errors.New("empty name")
	}

	return fmt.Sprintf(randomFormat(),name), nil;
}
func pHello(name string) {
	message, err := Hello(name)

	if err != nil{
		log.Fatal(err)
		return
	}
	fmt.Println(message);
}

func randomFormat() string {
	formats := []string{
		"Greetings, %v!",
		"Well! Well! Well! Look who the cat dragged in! If it isn't %v!",
		"Hail %v! It is a pleasure to meet your acquaintance!",
	}

	format := formats[rand.Intn(len(formats))];
	
	return format;
}

var bannedWords = []string{
 "hitler",  "satan", "nigger", "slave", "bitch", "slut", "rape","rapist","cunt","pussy","cp","pedophilia","pedophile","blowjob","doggy-style","cow-girl",
}

func censorBadWord(text string, showScore bool) string{
	words := strings.Fields(text)
	outstring := ""

	for i:=0; i < len(words); i++{
		currentWord := words[i]
		lower := strings.ToLower(currentWord)
		censored := false
		for j:=0; j < len(bannedWords); j++{
			bannedWord := bannedWords[j]

			
			if(strings.Contains(lower,bannedWord)||diffAnal(lower,bannedWord)>=6.3){

				stars := strings.Repeat("*",len(bannedWords[j])-1)
				if(showScore) {outstring = fmt.Sprintf("%v %v%v(%v)",outstring,string(bannedWord[0]),stars,diffAnal(lower,bannedWord));}else{outstring = fmt.Sprintf("%v %v%v",outstring,string(bannedWord[0]),stars);}
				censored = true
				break
			}


			if(j == len(bannedWords)-1 && !censored){
				if(showScore){outstring = fmt.Sprintf("%v %v(%v)",outstring,string(currentWord),diffAnal(lower,bannedWord))}else{outstring = fmt.Sprintf("%v %v",outstring,string(currentWord))}
			}
		}
	}
	return  outstring;
}

func diffAnal(word1 string, word2 string) float32{
	confi := 0.0
	avg_length := float64(len(word1) + len(word2))/2.0
	length_diff := math.Abs(float64(len(word1)-len(word2)))
	step_size := 10/float64(avg_length);
	for i:= 0; i<len(word1); i++{
		for j:= 0; j<len(word2);j++{
			if(string(word1[i])==string(word2[j])){
				pos := math.Abs(float64(j)-float64(i))
				confi += step_size*(1/(pos +1))*(1/(length_diff+1))
			}
		}
	}
	return float32(confi)
}