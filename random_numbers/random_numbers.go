package random_numbers

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Gera números aleatórios e salva no arquivo
func GenerateRandomNumbers(n int, filename string) error {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		num := rand.Intn(1000000) // Gera números entre 0 e 999999
		_, err := file.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadNumbers(filename string) ([]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var numbers []int
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers, nil
}
