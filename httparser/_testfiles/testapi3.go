package testapi

import (
	"fmt"
	"net/http"
)

// NotHandler3 does nothing
func NotHandler3(w http.ResponseWriter) {
	fmt.Println("do nothing again")
}
