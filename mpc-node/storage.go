package main

import (
	"os"
)

// SaveShare saves a node share
func SaveShare(data []byte) error {
	return os.WriteFile("share.json", data, 0600)
}

// LoadShare loads a node share
func LoadShare() ([]byte, error) {
	return os.ReadFile("share.json")
}
