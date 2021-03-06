package generator

import (
	//"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func Generate() string {
	rand.Seed(time.Now().UTC().UnixNano())
	domains := findAvailableDomainName()
	return domains
}

func findAvailableDomainName() string {
	consonants := []string{"h", "k", "h", "n", "m", "y", "d", "r"}
	vowels := []string{"a", "e", "i", "o", "u"}
	domainName := ""
	available := false

	for !available {
		domainName = constructString(consonants, vowels)
		//fmt.Print("testing " + domainName+".com" + " ... ")
		available = executeHostCommand(domainName + ".com")
		/*if available {
			fmt.Println("available !")
		} else {
			fmt.Println("not available.")
		}*/
	}

	return domainName
}

func constructString(array1 []string, array2 []string) string {
	s := []string{}
	for i := 0; i < 3; i++ {
		s = append(s, pickRandomLetter(array1))
		s = append(s, pickRandomLetter(array2))
	}
	str := strings.Join(s, "")
	return str
}

func pickRandomLetter(array []string) string {
	return array[randInt(1, len(array))]
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func executeHostCommand(s string) bool {
	cmd := exec.Command("host", s)
	_, err := cmd.Output()
	if err != nil {
		return true
	} else {
		return false
	}
}
