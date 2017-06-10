package totsuzen

import "fmt"

func ExampleToken_Totsuzenize() {
	result := NewToken("test").Totsuzenize().Text
	fmt.Println(result)
}
