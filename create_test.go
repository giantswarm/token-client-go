package client

import (
	"testing"
)

func TestCreate(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}
	userId := "test"
	_, err = c.Create(userId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateEmpty(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}
	userId := ""
	_, err = c.Create(userId)
	if err == nil {
		t.Fatal("Expected an error")
	} else if err.Error() != "Invalid arguments" {
		t.Fatalf("Expected invalid arguments, got %v", err)
	}
}

func testCreate(userId string) (*Token, error) {
	c, err := newTestClient()
	if err != nil {
		return nil, err
	}
	token, err := c.Create(userId)
	return &token, err
}
