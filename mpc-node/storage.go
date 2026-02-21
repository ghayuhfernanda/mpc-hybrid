package main

import (
    "encoding/json"
    "os"
)

func SavePartyData(filename string, data interface{}) error {
    out, _ := json.Marshal(data)
    return os.WriteFile(filename, out, 0600)
}

func LoadPartyData(filename string, v interface{}) error {
    raw, _ := os.ReadFile(filename)
    return json.Unmarshal(raw, v)
}