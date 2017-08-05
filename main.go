package main

import ("fmt"
		"io/ioutil"
		"strings"
		"os"
)

func main(){
	if len(os.Args)>1{
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil{
			fmt.Println("Error in filename")
		}else{
			str := string(file)
			cnt := strings.Count(str,"\n") + 1
			fmt.Println(cnt)
		}
	}else{
		fmt.Println("Please, enter filename path in command line")
	}
}