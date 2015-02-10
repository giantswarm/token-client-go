package client

import (
	"testing"

	apiSchemaPkg "github.com/catalyst-zero/api-schema"
)

func TestCreate(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}
	userId := "test"
	resp, err := apiSchemaPkg.FromHTTPResponse(c.Create(userId))
	if err != nil {
		t.Fatal(err)
	}
	var token Token
	if err := resp.UnmarshalData(&token); err != nil {
		t.Fatalf("Cannot unmarshal token %v", err)
	}
}

func TestCreateEmpty(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		t.Fatal(err)
	}
	userId := ""
	_, err = apiSchemaPkg.FromHTTPResponse(c.Create(userId))
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
	resp, err := apiSchemaPkg.FromHTTPResponse(c.Create(userId))
	if err != nil {
		return nil, err
	}
	token := &Token{}
	if err := resp.UnmarshalData(token); err != nil {
		return nil, err
	}
	return token, nil
}
