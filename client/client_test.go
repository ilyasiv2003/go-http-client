package client_test

import (
	"github.com/bozd4g/go-http-client/client"
	"github.com/mitchellh/mapstructure"
	"testing"
)

type Todo struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

func TestGetRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.Get("posts/10")

	t.Run("Returns a todo who have id as 10", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 10

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestGetRequestWithParameters (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.GetWithParameters("posts", Todo {
		Id: 11,
	})

	t.Run("Returns a todo who have id as 11", func(t *testing.T) {
		var got []Todo
		mapstructure.Decode(response.Data, &got)
		want := 11

		if len(got) < 1 || got[10].Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got[10].Id, want)
		}
	})
}

func TestPostRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.Post("posts/20")

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestPostRequestWithParameters (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.PostWithParameters("posts", Todo{
		Id: 21,
	})

	t.Run("Returns a todo who have id as 20", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 21

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPutRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.Put("posts/30")

	t.Run("Returns a todo who have id as 30", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 30

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPutRequestWithParameters (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.PutWithParameters("posts", Todo {
		Id: 31,
	})

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestDeleteRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.Delete("posts/40")

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestDeleteRequestWithParameters (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://jsonplaceholder.typicode.com/" }
	response := client.DeleteWithParameters("posts", Todo {
		Id: 41,
	})

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}