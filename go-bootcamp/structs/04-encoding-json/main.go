package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type permissions map[string]bool

// type user struct {
// 	Name        string      `json:"username"`
// 	Password    string      `json:"-"`
// 	Permissions permissions `json:"permissions,omitempty"`
// }

// // ENCODE TO JSON
// func main() {
// 	users := []user{
// 		{"inanc", "1234", nil},
// 		{"fab", "42", permissions{"admin": true}},
// 		{"dave", "644", permissions{"write": true}},
// 	}

// 	// json.Marshal(users)
// 	out, err := json.MarshalIndent(users, "", "\t")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(out))
// }

// DECODE FROM JSON

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"permissions"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user

	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		}

		fmt.Println()
	}

}
