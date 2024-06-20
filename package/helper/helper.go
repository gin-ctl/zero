package helper

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// BcryptHash Encrypt passwords using bcrypt.
func BcryptHash(password string) (pw string, err error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(fromPassword), err
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
	// Open the file for read and write operations.
	// If the file does not exist, an error is returned.
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Get the status information of the file.
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	// If the file is not empty, check if the last byte is a newline character.
	if stat.Size() > 0 {
		// Gets the last byte of the file.
		lastByte := make([]byte, 1)
		_, err = file.ReadAt(lastByte, stat.Size()-1)
		if err != nil {
			return fmt.Errorf("failed to read last byte: %w", err)
		}

		// If the last byte is not a newline character, a newline character is added.
		if lastByte[0] != '\n' {
			_, err = file.WriteString("\n")
			if err != nil {
				return fmt.Errorf("failed to write newline to file: %w", err)
			}
		}
	}

	// Append new content and wrap lines automatically.
	_, err = file.WriteString("\n" + content + "\n")
	if err != nil {
		return fmt.Errorf("failed to write content to file: %w", err)
	}

	err = file.Close()

	return nil
}

// GetFileContent get file content return string.
func GetFileContent(filePath string) (content string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	content = string(data)
	err = file.Close()
	return
}

// InsertImport 插入新的import路径到指定的Go文件中
func InsertImport(filePath, newImport, format, t string) error {
	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("无法读取文件: %w", err)
	}
	formatStr := format + "("
	// 找到import块
	startIndex := bytes.Index(data, []byte(formatStr))
	if startIndex == -1 {
		return fmt.Errorf("not found " + format)
	}
	endIndex := bytes.Index(data[startIndex:], []byte(")"))
	if endIndex == -1 {
		return fmt.Errorf("not found " + format)
	}
	endIndex += startIndex

	// 提取import块的内容并进行排序
	importBlock := data[startIndex:endIndex]
	importLines := strings.Split(string(importBlock), "\n")
	var imports []string
	for _, line := range importLines {
		line = strings.TrimSpace(line)
		if line != "" && line != formatStr && line != ")" {
			imports = append(imports, line)
		}
	}
	if !Contains(imports, newImport) {
		imports = append(imports, newImport)
	}
	sort.Strings(imports)

	// 构造新的import块
	var newImportBlock strings.Builder
	newImportBlock.WriteString(formatStr + "\n")
	for _, imp := range imports {
		newImportBlock.WriteString("\t" + imp + "\n")
	}
	newImportBlock.WriteString(t + ")")

	// 构造新的文件内容
	var newFileContent bytes.Buffer
	newFileContent.Write(data[:startIndex])
	newFileContent.WriteString(newImportBlock.String())
	newFileContent.Write(data[endIndex+1:])

	// 写入修改后的内容到文件
	err = os.WriteFile(filePath, newFileContent.Bytes(), os.ModePerm)
	if err != nil {
		return fmt.Errorf("无法写入文件: %w", err)
	}

	return nil
}

func GetModule(pwd string) (module string) {
	file, err := os.Open(pwd + "/go.mod")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}
	return
}

func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
