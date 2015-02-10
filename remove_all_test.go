package client

import (
	"testing"
)

func TestRemoveAll(t *testing.T) {
	userId := "test-user-1"
	token1, err := testCreate(userId)
	if err != nil {
		t.Fatal(err)
	}
	token2, err := testCreate(userId)
	if err != nil {
		t.Fatal(err)
	}

	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}

	// Remove all tokens for the user, should be ok
	_, err = c.RemoveAll(userId)
	if err != nil {
		t.Fatal(err)
	}

	t1, err := c.Get(token1.Id)
	if err == nil {
		t.Errorf("Expected a not found error got token %v", t1)
	}

	t2, err := c.Get(token2.Id)
	if err == nil {
		t.Errorf("Expected a not found error, got token %v", t2)
	}
}

func TestRemoveAllNotFound(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}

	// Remove the token, should be ok
	_, err = c.RemoveAll("non-existing-user-id")
	if err != nil {
		t.Fatal(err)
	}
}
