package envfile

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var Debug bool

func EnvFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("can't open %q -- %w", filename, err)
	}
	defer f.Close()
	return EnvLoad(f)
}

func EnvLoad(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		if strings.HasPrefix(text, "#") {
			continue
		}
		if strings.HasPrefix(text, "export ") {
			const skip = len("export ")
			text = text[skip:]
			text = strings.TrimSpace(text)
		}
		key, value, ok := strings.Cut(text, "=")
		if ok {
			value = os.ExpandEnv(value)
			value = strings.Trim(value, `"`)
			if Debug {
				log.Printf("Setting env %q to %q", key, value)
			}
			os.Setenv(key, value)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}
	return nil
}
