package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/BlokOfWood/CSK4-szamtan-gyakorlo/backend/internal/assert"
	"github.com/BlokOfWood/CSK4-szamtan-gyakorlo/backend/internal/data"
)

func TestCreateResult(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	t.Run("valid input", func(t *testing.T) {
		jsonStr := `{"name":"test","score":100}`
		statusCode, _, body := ts.post(t, "/v1/results", jsonStr)

		assert.Equal(t, statusCode, http.StatusCreated)

		var wrapper struct {
			Result *data.Result `json:"result"`
		}
		err := json.Unmarshal([]byte(body), &wrapper)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, wrapper.Result.Name, "test")
		assert.Equal(t, wrapper.Result.Score, 100)
	})

	t.Run("invalid input", func(t *testing.T) {
		jsonStr := `{"name":"test","score":1000,"extra":"field"}`
		statusCode, _, _ := ts.post(t, "/v1/results", jsonStr)

		assert.Equal(t, statusCode, http.StatusBadRequest)
	})
}

func TestShowResult(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	t.Run("existing record", func(t *testing.T) {
		result := &data.Result{
			Name:  "test",
			Score: 100,
		}
		err := app.models.Results.Insert(result)
		if err != nil {
			t.Fatal(err)
		}

		statusCode, _, body := ts.get(t, "/v1/results/1")

		assert.Equal(t, statusCode, http.StatusOK)

		var wrapper struct {
			Result *data.Result `json:"result"`
		}
		err = json.Unmarshal([]byte(body), &wrapper)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, wrapper.Result.ID, result.ID)
		assert.Equal(t, wrapper.Result.Name, result.Name)
		assert.Equal(t, wrapper.Result.Score, result.Score)
	})

	t.Run("record not found", func(t *testing.T) {
		statusCode, _, _ := ts.get(t, "/v1/results/2")

		assert.Equal(t, statusCode, http.StatusNotFound)
	})
}

func TestListResults(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	t.Run("no records", func(t *testing.T) {
		statusCode, _, body := ts.get(t, "/v1/results")

		assert.Equal(t, statusCode, http.StatusOK)

		var wrapper struct {
			Results []*data.Result `json:"results"`
		}
		err := json.Unmarshal([]byte(body), &wrapper)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(wrapper.Results), 0)
	})

	t.Run("multiple records", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			result := &data.Result{
				Name:  "test",
				Score: 100,
			}
			err := app.models.Results.Insert(result)
			if err != nil {
				t.Fatal(err)
			}
		}

		statusCode, _, body := ts.get(t, "/v1/results")

		assert.Equal(t, statusCode, http.StatusOK)

		var wrapper struct {
			Results []*data.Result `json:"results"`
		}
		err := json.Unmarshal([]byte(body), &wrapper)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(wrapper.Results), 5)
	})
}
