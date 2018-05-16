package httpflags_test

import (
	"fmt"
	"net/http/httptest"

	"github.com/artyom/httpflags"
)

func ExampleParse() {
	args := &struct {
		Name  string `flag:"name"`
		Age   int    `flag:"age"`
		Extra bool   `flag:"extra"`
	}{
		// default values:
		Age:   42,
		Extra: true,
	}
	req := httptest.NewRequest("GET", "/?name=John%20Doe&extra=false", nil)
	if err := httpflags.Parse(args, req); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("updated args: %+v\n", args)

	req = httptest.NewRequest("GET", "/?badField=boom", nil)
	fmt.Println("parsing request with undefined field:", httpflags.Parse(args, req))
	fmt.Printf("args stay the same: %+v\n", args)
	// Output:
	// updated args: &{Name:John Doe Age:42 Extra:false}
	// parsing request with undefined field: <nil>
	// args stay the same: &{Name:John Doe Age:42 Extra:false}
}
