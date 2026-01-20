package storage

import "main/internal/model"

func (s *Storage) Migrate() {
	s.DB.AutoMigrate(&model.Number{})
}
