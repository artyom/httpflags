package httpflags_test

import (
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"

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

func TestParse(t *testing.T) {
	args := struct {
		IDs idList `flag:"id"`
	}{}
	r := httptest.NewRequest("GET", "/?id=1&id=2&id=3", nil)
	if err := httpflags.Parse(&args, r); err != nil {
		t.Fatal(err)
	}
	if want := 3; len(args.IDs) != want {
		t.Fatalf("want %d parsed values, got: %v", want, args.IDs)
	}
}

type idList []uint64

func (l *idList) String() string { return "n/a" }
func (l *idList) Set(value string) error {
	id, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return err
	}
	*l = append(*l, id)
	return nil
}
