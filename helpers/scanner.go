package helpers

import (
	"bufio"
	"os"
)

func Scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
