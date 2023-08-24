package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"

	"github.com/nfjBill/gorm-driver-dm/dmr"
)

// TODO: ignore
// StringToDMClob convert string to dmr.DmClob
func StringToDMClob(s string) dmr.DmClob {
	return *dmr.NewClob(s)
}

// CompressText compress text with gzip
func CompressText(text string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write([]byte(text)); err != nil {
		fmt.Println("Compression error:", err)
		return nil, err
	}
	if err := gz.Close(); err != nil {
		fmt.Println("Compression error:", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecompressText decompress text with gzip
func DecompressText(compressed []byte) (string, error) {
	buf := bytes.NewReader(compressed)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		fmt.Println("Decompression error:", err)
		return "", err
	}
	defer gz.Close()
	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(gz); err != nil {
		fmt.Println("Decompression error:", err)
		return "", err
	}
	return string(buffer.Bytes()), nil
}

// Base64Encode encode bytes to base64 string
func Base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// Base64Decode decode base64 string to bytes
func Base64Decode(text string) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return []byte(""), err
	}

	return []byte(decodedData), nil
}

// CompressAndEncodeBase64 compress and encode json data to base64 string
func CompressAndEncodeBase64(jsonData string) (string, error) {
	data := []byte(jsonData)

	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	encodedString := base64.StdEncoding.EncodeToString(buf.Bytes())

	return encodedString, nil
}

// DecodeAndDecompressBase64 decode and decompress base64 string to json data
func DecodeAndDecompressBase64(encodedString string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	buf.Write(decodedData)

	gz, err := gzip.NewReader(&buf)
	if err != nil {
		return "", err
	}

	decompressedData := new(bytes.Buffer)
	if _, err := decompressedData.ReadFrom(gz); err != nil {
		return "", err
	}

	if err := gz.Close(); err != nil {
		return "", err
	}

	result := decompressedData.String()

	return result, nil
}
