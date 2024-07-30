package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	mathrand "math/rand"
	"strings"
	"time"
)

const lettersAndDigits = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const specialCharacters = "@#$%&!" // add more symbols as needed: *, ', ...etc

// Generate a strong password with groups and separators
func generateStrongPassword(groupLength, numGroups, numSpecials int, separator string, rng *mathrand.Rand) string {
	// Generate special characters for the entire password
	specials := make([]rune, numSpecials)
	for i := range specials {
		char, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialCharacters))))
		specials[i] = rune(specialCharacters[char.Int64()])
	}

	// Generate the remaining characters
	remainingChars := make([]rune, groupLength*numGroups-numSpecials)
	for i := range remainingChars {
		char, _ := rand.Int(rand.Reader, big.NewInt(int64(len(lettersAndDigits))))
		remainingChars[i] = rune(lettersAndDigits[char.Int64()])
	}

	// Combine all characters and shuffle
	allChars := append(specials, remainingChars...)
	rng.Shuffle(len(allChars), func(i, j int) {
		allChars[i], allChars[j] = allChars[j], allChars[i]
	})

	// Split into groups
	var passwordGroups []string
	for i := 0; i < len(allChars); i += groupLength {
		if i+groupLength <= len(allChars) {
			passwordGroups = append(passwordGroups, string(allChars[i:i+groupLength]))
		} else {
			passwordGroups = append(passwordGroups, string(allChars[i:]))
		}
	}

	// Join groups with separator
	return strings.Join(passwordGroups, separator)
}

func main() {
	numPasswords := flag.Int("n", 1, "Number of passwords to generate. Default is 1.")
	totalLength := flag.Int("l", 16, "Total length of each password, including separators. Default is 16.")
	groupLength := flag.Int("g", 4, "Length of each group, excluding separators. Default is 4.")
	separator := flag.String("s", "-", "Separator between groups. Default is '-'.")
	numSpecials := flag.Int("c", 2, "Number of special characters in each password. Default is 2.")

	flag.Parse()

	if *totalLength < *numSpecials {
		log.Fatal("Total length must be at least equal to the number of special characters.")
	}

	// Calculate number of groups and separators
	numGroups := (*totalLength + len(*separator)) / (*groupLength + len(*separator))
	actualLength := numGroups*(*groupLength) + (numGroups-1)*len(*separator)

	// Ensure the calculated length fits within the total length
	if actualLength > *totalLength {
		log.Fatalf("Total length (%d) is not sufficient for the specified group length (%d), number of groups (%d), and separators.",
			*totalLength, *groupLength, numGroups)
	}

	// Create a new local random generator
	rng := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *numPasswords; i++ {
		password := generateStrongPassword(*groupLength, numGroups, *numSpecials, *separator, rng)
		fmt.Println("Generated Password:", password)
	}
}
