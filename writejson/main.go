package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type song struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

var name, aut, op string

func savePlaylist(playlist []song) {
	fi, err := os.Create("my_playlist.json") // read access
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(fi).Encode(playlist)
	defer fi.Close()
}

func LoadPlaylist(playlist *[]song) {
	fi, err := os.Open("my_playlist.json") // read access
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(fi).Decode(playlist)
	defer fi.Close()
}

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
	defer fmt.Println("========== END OF SONG LIST ==========")
	for i, s := range playlist {
		fmt.Printf("%d. %s by %s\n", i+1, s.Name, s.Author)
	}
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
		} else if op == "save" {
			savePlaylist(*playlist)
		} else if op == "load" {
			LoadPlaylist(playlist)
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
