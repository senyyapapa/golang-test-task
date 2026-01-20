package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"log/slog"
	"main/internal/model"
	"main/internal/storage"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestApp(t *testing.T) (*fiber.App, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.Number{})
	require.NoError(t, err)

	store := &storage.Storage{DB: db}
	logger := slog.Default()
	handler := NewHandler(store, logger)

	app := fiber.New()
	app.Post("/number", handler.PostNum)
	app.Get("/number", handler.GetNumber)

	return app, func() {}
}

func TestPostAndGetNumbers(t *testing.T) {
	app, cleanup := newTestApp(t)
	defer cleanup()

	// POST /number
	body := bytes.NewBuffer([]byte(`{"Num": 100}`))
	req := httptest.NewRequest("POST", "/number", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	// GET /number
	req = httptest.NewRequest("GET", "/number", nil)
	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var result []int
	err = json.NewDecoder(resp.Body).Decode(&result)
	require.NoError(t, err)
	assert.Equal(t, []int{100}, result)
}
