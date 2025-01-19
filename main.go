package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var maxNotes int
	for {
		fmt.Println("Enter the maximum number of notes:")
		_, err := fmt.Scan(&maxNotes)

		if err != nil {
			fmt.Println("[Error] Invalid input, not a number.\n")

			var discard string
			fmt.Scanln(&discard)
		} else {
			break
		}
	}

	notes := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nEnter a command and data:")
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		switch command {
		case "create":
			create(&notes, maxNotes, data)
		case "list":
			list(notes)
		case "clear":
			clear(&notes)
		case "update":
			update(&notes, input, maxNotes, command)
		case "delete":
			deleteNote(&notes, input, maxNotes, command)
		case "exit":
			exit()
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func create(notes *[]string, maxNotes int, data string) {
	data = strings.TrimSpace(data)

	if len(*notes) < maxNotes {
		if data == "" {
			fmt.Println("[Error] Missing note argument")
		} else {
			*notes = append(*notes, data)
			fmt.Println("[OK] The note was successfully created")
		}
	} else {
		fmt.Println("[Error] Notepad is full")
	}
}

func list(notes []string) {
	if len(notes) == 0 {
		fmt.Println("[Info] Notepad is empty")
	} else {
		for i, note := range notes {
			fmt.Printf("[Info] %d: %s\n", i+1, note)
		}
	}
}

func update(notes *[]string, input []string, maxNotes int, command string) {
	position, err := checkInput(notes, input, maxNotes, command, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	note := strings.Join(input[2:], " ")
	(*notes)[position-1] = note
	fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
}

func clear(notes *[]string) {
	*notes = []string{}
	fmt.Println("[OK] All notes were successfully deleted")
}

func deleteNote(notes *[]string, input []string, maxNotes int, command string) {
	position, err := checkInput(notes, input, maxNotes, command, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	*notes = append((*notes)[:position-1], (*notes)[position:]...)
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
}

func exit() {
	fmt.Println("[Info] Bye!")
	os.Exit(0)
}

func checkInput(notes *[]string, input []string, maxNotes int, command string, requiresNote bool) (int, error) {
	if len(input) < 2 {
		return 0, fmt.Errorf("[Error] Missing position argument")
	}

	position, err := strconv.Atoi(input[1])
	if err != nil {
		return 0, fmt.Errorf("[Error] Invalid position: %s", input[1])
	}

	if requiresNote && len(input) < 3 {
		return 0, fmt.Errorf("[Error] Missing note argument")
	}

	if position < 1 || position > maxNotes {
		return 0, fmt.Errorf("[Error] Position %d is out of the boundaries [1, %d]", position, maxNotes)
	}

	if position > len(*notes) && position <= maxNotes {
		if command == "update" {
			return 0, fmt.Errorf("[Error] There is nothing to update")
		} else {
			return 0, fmt.Errorf("[Error] There is nothing to delete")
		}
	}

	return position, nil
}
