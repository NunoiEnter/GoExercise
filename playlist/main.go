package main

import (
	"fmt"
)

type song struct {
	Name   string
	Author string
}

var name, aut, op string

func addSong(playlist *[]song) []song {
	fmt.Println("Enter Song Name : ")
	fmt.Scanln(&name)
	fmt.Println("Enter Song Author : ")
	fmt.Scanln(&aut)
	newList := song{Name: name, Author: aut}
	*playlist = append(*playlist, newList)
	return *playlist
}

func rmSong(playlist *[]song) []song {
	fmt.Println("Enter Remove song : ")
	fmt.Scanln(&name)
	windex := 0
	for _, s := range *playlist {
		if s.Name != name {
			(*playlist)[windex] = s
			windex++
		}
	}
	*playlist = (*playlist)[:windex]
	return *playlist
}

func lsSong(playlist []song) {
	fmt.Println("========== LIST OF SONG ==========")
	for i, s := range playlist {
		fmt.Printf("%d. %s by %s\n", i+1, s.Name, s.Author)
	}
	fmt.Println("========== END OF SONG LIST ==========")
}

func Operation(playlist *[]song) {
	cnt := 0
	for {
		fmt.Println("Enter Your Operation (add/rm/ls/exit)")
		fmt.Scanln(&op)
		if op == "exit" {
			break
		} else if op == "add" {
			addSong(playlist)
		} else if op == "rm" {
			rmSong(playlist)
		} else if op == "ls" {
			lsSong(*playlist)
		} else {
			fmt.Println("There is no such a Operation Try Again")
			cnt += 1
			if cnt == 3 {
				fmt.Println("Your Type Skill is ass Session Terminate")
				break
			}
		}
	}
}

func main() {
	playlist := []song{}
	Operation(&playlist)
}
