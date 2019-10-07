package b2048

import (
    "errors"
    "strings"
)

// Encoding is a custom base encoding defined by an alphabet.
// It should bre created using NewEncoding function
type Encoding struct {
    base        int
    alphabet    []string
    alphabetMap map[string]int
}

var (
    e *Encoding
)

// NewEncoding returns a custom base encoder defined by the alphabet string.
// The alphabet should contain non-repeating characters.
// Ordering is important.
// Example alphabets:
//   - base2: 01
//   - base16: 0123456789abcdef
//   - base32: 0123456789ABCDEFGHJKMNPQRSTVWXYZ
//   - base62: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
func NewEncoding(alphabet []string) (*Encoding, error) {
    runes := alphabet
    runeMap := make(map[string]int)

    for i := 0; i < len(runes); i++ {
        if _, ok := runeMap[runes[i]]; ok {
            return nil, errors.New("ambiguous alphabet")
        }

        runeMap[runes[i]] = i
    }

    return &Encoding{
        base:        len(runes),
        alphabet:    runes,
        alphabetMap: runeMap,
    }, nil
}

func init() {
    e, _ = NewEncoding(WorldList)
}

// Encode function receives a byte slice and encodes it to a string using the alphabet provided
func Encode(source []byte) string {
    if len(source) == 0 {
        return ""
    }

    digits := []int{0}

    for i := 0; i < len(source); i++ {
        carry := int(source[i])

        for j := 0; j < len(digits); j++ {
            carry += digits[j] << 8
            digits[j] = carry % e.base
            carry = carry / e.base
        }

        for carry > 0 {
            digits = append(digits, carry%e.base)
            carry = carry / e.base
        }
    }

    var res []string

    for k := 0; source[k] == 0 && k < len(source)-1; k++ {
        res = append(res, e.alphabet[0])
    }

    for q := len(digits) - 1; q >= 0; q-- {
        res = append(res, e.alphabet[digits[q]])
    }

    return strings.Join(res, "-")
}

// Decode function decodes a string previously obtained from Encode, using the same alphabet and returns a byte slice
// In case the input is not valid an error will be returned
func Decode(source string) ([]byte, error) {
    runes := strings.Split(source, "-")

    if len(runes) == 0 {
        return []byte{}, nil
    }

    returnBytes := []byte{0}
    for i := 0; i < len(runes); i++ {
        value, ok := e.alphabetMap[runes[i]]

        if !ok {
            return nil, errors.New("Non Base Character")
        }

        carry := int(value)

        for j := 0; j < len(returnBytes); j++ {
            carry += int(returnBytes[j]) * e.base
            returnBytes[j] = byte(carry & 0xff)
            carry >>= 8
        }

        for carry > 0 {
            returnBytes = append(returnBytes, byte(carry&0xff))
            carry >>= 8
        }
    }

    for k := 0; runes[k] == e.alphabet[0] && k < len(runes)-1; k++ {
        returnBytes = append(returnBytes, 0)
    }

    // Reverse returnBytes
    for i, j := 0, len(returnBytes)-1; i < j; i, j = i+1, j-1 {
        returnBytes[i], returnBytes[j] = returnBytes[j], returnBytes[i]
    }

    return returnBytes, nil
}