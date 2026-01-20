package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main/internal/model"
)

// Вспомогательная функция только для тестов
func newTestStorage(t *testing.T) (*Storage, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.Number{})
	require.NoError(t, err)

	return &Storage{DB: db}, func() {}
}

func TestPostNumAndGetArrayNum(t *testing.T) {
	s, cleanup := newTestStorage(t)
	defer cleanup()

	// Тестируем PostNum
	err := s.PostNum(42)
	assert.NoError(t, err)

	err = s.PostNum(7)
	assert.NoError(t, err)

	// Тестируем GetArrayNum
	nums := s.GetArrayNum()
	assert.Equal(t, []int{7, 42}, nums) // сортировка по Num ASC
}
