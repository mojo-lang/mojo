package _go

import (
	"strings"
	"testing"
)

func TestDiffGoCode(t *testing.T) {
	code := []string{
		// Normal
		`func Foo() {
			println("test_data")
		}`,
		// Leading and trailing whitespace
		`
		
		func Foo () {
			println("test_data")
		}
		
		`,
		// Indentation differences, tabs
		`				func Foo() {
										println("test_data")
			}`,
		// Indentation differences, spaces
		`    func Foo() {
		                                                    println("test_data")
				    				            }
	                  `,
	}

	for _, v := range code {
		a, b, di := DiffGoCode(code[0], v)
		if strings.Compare(a, b) != 0 {
			t.Errorf("Code differs: %s", di)
		}
	}
}
