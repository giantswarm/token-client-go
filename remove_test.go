package client

import (
	"testing"
)

func TestRemove(t *testing.T) {
	userId := "test-user-1"
	token, err := testCreate(userId)
	if err != nil {
		t.Fatal(err)
	}

	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}

	// Remove the token, should be ok
	_, err = c.Remove(token.Id)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Get(token.Id)
	if err == nil {
		t.Fatal("Expected a not found error")
	}
}

func TestRemoveNotFound(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}

	// Remove the token, should be ok
	_, err = c.Remove("non-existing-token-id")
	if err != nil {
		t.Fatal(err)
	}
}
