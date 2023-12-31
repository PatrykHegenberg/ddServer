package model

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Monster struct für die Daten des Monsters
type Monster struct {
	Save            Save     `json:"save"`
	Skill           Skill    `json:"skill"`
	HP              HP       `json:"hp"`
	Source          string   `json:"source"`
	CR              string   `json:"cr"`
	Type            string   `json:"type"`
	Name            string   `json:"name"`
	Vulnerable      []string `json:"vulnerable"`
	ConditionImmune []string `json:"conditionImmune"`
	Resist          []string `json:"resist"`
	Immune          []string `json:"immune"`
	Traits          []Trait  `json:"trait"`
	AC              []AC     `json:"ac"`
	Alignment       []string `json:"alignment"`
	Senses          []string `json:"senses"`
	Languages       []string `json:"languages"`
	Size            []string `json:"size"`
	Actions         []Action `json:"action"`
	Speed           Speed    `json:"speed"`
	Str             int      `json:"str"`
	Dex             int      `json:"dex"`
	Con             int      `json:"con"`
	Int             int      `json:"int"`
	Wis             int      `json:"wis"`
	Cha             int      `json:"cha"`
}

type AC struct {
	From []string `json:"from"`
	AC   int      `json:"ac"`
}

type HP struct {
	Formula string `json:"formula"`
	Average int    `json:"average"`
}

type Speed struct {
	Walk   int `json:"walk"`
	Burrow int `json:"burrow"`
	Climb  int `json:"climb"`
	Fly    int `json:"fly"`
	Swim   int `json:"swim"`
}

type Save struct {
	Dex string `json:"dex"`
	Con string `json:"con"`
	Wis string `json:"wis"`
	Cha string `json:"cha"`
	Str string `json:"str"`
	Int string `json:"int"`
}

type Skill struct {
	Stealth        string `json:"stealth"`
	Acrobatics     string `json:"acrobatics"`
	AnimalHandling string `json:"animalHandling"`
	Arcana         string `json:"arcana"`
	Athletics      string `json:"athletics"`
	Deception      string `json:"deception"`
	History        string `json:"history"`
	Insight        string `json:"insight"`
	Intimidation   string `json:"intimidation"`
	Investigation  string `json:"investigation"`
	Medicine       string `json:"medicine"`
	Nature         string `json:"nature"`
	Perception     string `json:"perception"`
	Performance    string `json:"performance"`
	Persuasion     string `json:"persuasion"`
	SleightOfHand  string `json:"sleightOfHand"`
	Survival       string `json:"survival"`
	Religion       string `json:"religion"`
}

type Trait struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

type Action struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

// Character struct für die Daten des Charakters
type Character struct {
	Meta    Meta      `json:"_meta"`
	Monster []Monster `json:"monster"`
}

// Meta struct für Meta-Informationen
type Meta struct {
	DateLastModifiedHash string   `json:"_dateLastModifiedHash"`
	Sources              []Source `json:"sources"`
	DateAdded            int64    `json:"dateAdded"`
	DateLastModified     int64    `json:"dateLastModified"`
}

type Source struct {
	Json         string   `json:"json"`
	Abbreviation string   `json:"abbreviation"`
	Version      string   `json:"version"`
	Authors      []string `json:"authors"`
	ConvertedBy  []string `json:"convertedBy"`
}

// WriteToFile writes data to a file.
// It takes in a filename string and a data byte slice.
// It returns an error if there was an issue writing to the file, otherwise it returns nil.
func WriteToFile(filename string, data []byte) error {
	log.Println("Writing data to file:", filename)

	// Create a file with the given filename
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing file:", err)
		}
	}()

	// Write the data to the file
	n, err := file.Write(data)
	if err != nil {
		log.Println("Error writing to file:", err)
		return err
	}
	log.Printf("Successfully wrote %d bytes to file", n)

	return nil
}

// getOrCreateCharacter returns the current character object or creates a new one
func GetOrCreateCharacter(filename string, chars []Character) Character {
	// Check if there is an empty character object
	for _, char := range chars {
		if char.Meta.DateLastModified == 0 {
			// Return the empty character object
			log.Println("Returning existing character object")
			return char
		}
	}

	// Create a new character object
	now := time.Now().Unix()
	newChar := Character{
		Meta: Meta{
			Sources: []Source{
				{
					Json:         "Malgorgon",
					Abbreviation: "MG",
					Authors:      []string{"Krzysztof"},
					ConvertedBy:  []string{"Krzysztof"},
					Version:      "unknown",
				},
			},
			DateAdded:            now,
			DateLastModified:     now,
			DateLastModifiedHash: fmt.Sprintf("%x", now),
		},
		Monster: []Monster{},
	}

	// Append the new character object to the list of characters
	chars = append(chars, newChar)

	// Return the newly created character object
	log.Println("Returning newly created character object")
	return newChar
}
