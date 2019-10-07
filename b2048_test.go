package b2048

import (
    "fmt"
    "testing"
)

func Test_Encode(t *testing.T) {
    phrase := "bd843"

    encode := Encode([]byte(phrase))
    fmt.Println("Encoded:", encode)

    decoded, err := Decode(encode)
    fmt.Println("Decoded:", string(decoded))
    if err != nil {
        t.Fatal("Failed to decode")
    }

    if phrase != string(decoded) {
        t.Fatal("Not the same, got:", string(decoded))
    }

    t.Log("Encoded:", encode)
}
