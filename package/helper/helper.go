package helper

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// BcryptHash Encrypt passwords using bcrypt.
func BcryptHash(password string) (pw string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// BcryptCheck Compare the plaintext password with the database hash.
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Empty Whether the detection value exists.
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr Converts time to a string of milliseconds.
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func CamelCase(input string) string {
	var output []rune
	toUpper := true

	for _, r := range input {
		if r == '_' {
			toUpper = true
			continue
		}
		if toUpper {
			output = append(output, unicode.ToUpper(r))
			toUpper = false
		} else {
			output = append(output, unicode.ToLower(r))
		}
	}

	return string(output)
}

// Capitalize Capital case
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	// read first rune
	r, size := utf8.DecodeRuneInString(s)
	// Convert the first rune to uppercase and concatenate the rest
	return string(unicode.ToUpper(r)) + s[size:]
}

// CreateDirIfNotExist Create dir if not existed.
func CreateDirIfNotExist(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

// PathExists checks if the specified path (file or directory) exists and returns a boolean value.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReadLines Reads file contents into string slices.
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteLines Write the modified content back to the file.
func WriteLines(filePath string, lines []string) error {
	content := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}

// CheckLineIsExisted Check str is existed.
func CheckLineIsExisted(lines []string, new string) bool {
	for _, line := range lines {
		if strings.Contains(line, new) {
			return true
		}
	}
	return false
}

// InsertOffset Insert the invoke call at the specified location.
func InsertOffset(lines []string, new, offset string) []string {
	for i, line := range lines {
		if strings.Contains(line, offset) {
			lines[i] = new + "\n" + line
			break
		}
	}
	return lines
}

// AppendToFile appends the given content to the end of the specified file
func AppendToFile(filePath string, content string) error {
	// Read the existing content from the file
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Append the new content to the existing content
	newContent := append(existingContent, content...)

	// Write the new content back to the file
	err = os.WriteFile(filePath, newContent, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// GetFileContent get file content return string.
func GetFileContent(filePath string) (content string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	content = string(data)
	return
}
