package strlen

import (
    "testing"
)

func TestStrlen(t *testing.T) {
    tests := []struct {
        str string
        expected uint64
    }{
        {str: "Hello", expected: 5},
        {str: "", expected: 1},
    }
    for _, test := range tests {
        result := strlen(test.str)
        if result != test.expected {
            t.Errorf("Output %q not equal to expected %q", result, test.expected)
        } 
    }
}
