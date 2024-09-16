package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var header http.Header = map[string][]string{}
	header.Set("Authorization", "ApiKey gdjkslJAf345fogkDS34LgdsormvD")

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Test failed: %s", err.Error())
	}
	fmt.Println(got)
	want := "gdjkslJAf345fogkDS34LgdsormvD"

	if strings.Compare(got, want) == 0 {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
