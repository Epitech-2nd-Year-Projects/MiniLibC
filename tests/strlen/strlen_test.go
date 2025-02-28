package strlen

import (
    "testing"
)

func TestStrlen(t *testing.T) {
    tests := []struct {
        str string
        expected uint
    }{
        {str: "Hello", expected: 5},
        {str: "", expected: 0},
        {str: " Hello ", expected: 7},
    }
    for _, test := range tests {
        result := strlen(test.str)
        if result != test.expected {
            t.Errorf("Output %d not equal to expected %d", result, test.expected)
        } 
    }
}
