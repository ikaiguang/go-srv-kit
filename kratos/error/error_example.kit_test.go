package errorpkg

import (
	"fmt"
	"net/http"
)

func ExampleNew() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	fmt.Printf("%+v\n", err)
}

func ExampleCause() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	fmt.Println(Cause(err))
}

func ExampleFromError() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	cause := FromError(err)
	fmt.Println("error Code ", cause.Code)
	fmt.Println("error Reason ", cause.Reason)
	fmt.Println("error Message ", cause.Message)
	fmt.Println("error Metadata ", cause.Metadata)
}

func ExampleIsCode() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	fmt.Println(IsCode(err, http.StatusNotFound))
}

func ExampleIsReason() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	fmt.Println(IsReason(err, "StatusNotFound"))
}

func ExampleCode() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	fmt.Println(Code(err))
}

func ExampleReason() {
	err := New(http.StatusNotFound, "StatusNotFound", "StatusNotFound")
	fmt.Println(Reason(err))
}
