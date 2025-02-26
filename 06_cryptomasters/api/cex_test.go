package api_test

import "testing"

func TestAPICall(t *testing.T) {
    _, err := api.getRate("")
    if err == nil {
        t.Error("Error was not found")
    }
}
