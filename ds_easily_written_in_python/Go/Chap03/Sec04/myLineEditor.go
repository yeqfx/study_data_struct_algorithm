package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"../../dataStruct/ArrayList"
)

func main() {
	list := &ArrayList.Arraylist{}
LoopFor:
	for {
		var command string
		fmt.Print("[메뉴선택] i-입력, d-삭제, r-변경, p-출력, l-파일읽기, s-저장, q-종료=> ")
		_, err := fmt.Scanln(&command)
		if err != nil {
			log.Fatal(err)
		}
		switch command {
		case "i":
			var pos int
			var str string
			fmt.Print("  입력행 번호 : ")
			fmt.Scanln(&pos)
			fmt.Print("  입력행 내용 : ")
			reader := bufio.NewReader(os.Stdin)
			str, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			// _, err := fmt.Scanln(&str)
			// if err != nil {
			// 	fmt.Println(err)
			// 	stdin.ReadString('\n')
			// }
			list = list.Insert(pos, str)

		case "d":
			var pos int
			fmt.Print("  삭제행 번호 : ")
			fmt.Scanln(&pos)
			list.Delete(pos)

		case "r":
			var pos int
			var str string
			fmt.Print("  변경행 번호 : ")
			fmt.Scanln(&pos)
			fmt.Print("  변경행 내용 : ")
			fmt.Scanln(&str)
			list.Replace(pos, str)

		case "q":
			break LoopFor

		case "p":
			fmt.Println("Line Editor")
			for pos, str := range *list {
				fmt.Printf("[%2d] %s", pos, str)
			}
			fmt.Println("")

		case "l":
			file, err := os.Open("test.txt")
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			scanner := bufio.NewScanner(file)
			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
			}
			for i := 0; scanner.Scan(); i++ {
				fmt.Printf("[%2d] %s\n", i, scanner.Text())
			}

		case "s":
			file, err := os.Create("test.txt")
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			for _, line := range *list {
				fmt.Fprint(file, fmt.Sprint(line))
			}
		}
	}
}
