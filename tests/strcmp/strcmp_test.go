package strcmp

import (
	"testing"
)

func TestStrcmp(t *testing.T) {
    tests := []struct {
        str1 string
        str2 string
        expected int
    }{
        {str1: "coucou", str2: "coucou", expected: 0},
        {str1: "abcd", str2: "abCd", expected: 32},
    }
    for _, test := range tests {
        result := strcmp(test.str1, test.str2)
        t.Log(result)
//        if convert != test.expected {
  //          t.Errorf("Ouput %d not equal to expected %d", result, test.expected)
    //    }
    }
}
