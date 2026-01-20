package storage

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main/internal/model"
	"os"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage() (*Storage, error) {
	db, err := gorm.Open(sqlite.Open("solution.db"))
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) PostNum(num int) error {
	newNum := &model.Number{
		Num: num,
	}

	res := s.DB.Create(newNum)

	if res.Error != nil {
		return res.Error
	}

	fmt.Fprintln(os.Stdout, "Number added to database")
	return nil
}

func (s *Storage) GetArrayNum() []int {
	var numbers []model.Number
	s.DB.Order("Num").Find(&numbers)

	res := make([]int, len(numbers))
	for i, n := range numbers {
		res[i] = n.Num
	}
	return res
}
