package client

import (
	"testing"
)

func TestGet(t *testing.T) {
	userId := "test-user-1"
	token, err := testCreate(userId)
	if err != nil {
		t.Fatal(err)
	}

	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}

	token2, err := c.Get(token.Id)
	if err != nil {
		t.Fatal(err)
	}

	if token2 != *token {
		t.Fatalf("Got different token. Expected %v, got %v", *token, token2)
	}
}

func TestGetNotFound(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Get("invalid-token-id")
	if err == nil {
		t.Fatal("Expected to get an error")
		/*} else if !IsErrTokenNotFound(err) {
		// Enable this after the error-fix in tokend
		t.Fatalf("Expected a not-found error, got %v", err)*/
	}
}
