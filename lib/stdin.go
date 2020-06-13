package lib

import (
	"bufio"
	"os"
)

var (
	stdin *bufio.Scanner
)

// InitStdin 標準入力を読むScannerの初期化
func InitStdin() {
	stdin = bufio.NewScanner(os.Stdin)
}

// ReadLine 標準入力から一行読んで返す
func ReadLine() string {
	stdin.Scan()
	return stdin.Text()
}
