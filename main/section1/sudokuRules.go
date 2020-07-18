package main

import (
	"errors"
	"fmt"
	"os"
)

type SudokuField [9][9]int8

type Sudoku struct {
	original, modified *SudokuField
}

func NewSudoku(field *SudokuField) (*Sudoku, error) {
	if field == nil {
		return nil, errors.New("field should point to something")
	}
	return &Sudoku{original: field, modified: new(SudokuField)}, nil
}

func (s *Sudoku) Set(x, y, number int8) error {
	if s == nil {
		return errors.New("can not set value to empty sudoku")
	}
	if isNumberNotOnField(x+1) || isNumberNotOnField(y+1) || isNumberNotOnField(number) {
		return errors.New("diapason of field is out of bounds")
	}
	if err := s.check(x, y, number); err != nil {
		return err
	}
	s.modified[x][y] = number
	return nil
}

func (s *Sudoku) check(x, y, number int8) error {
	value := s.modified[x][y]
	s.modified[x][y] = 0
	defer func() { s.modified[x][y] = value }()
	if !(s.original[x][y] == 0 && s.checkVertical(y, number) && s.checkGrid(x, y, number) &&
		s.checkHorizontal(x, number)) {
		return errors.New("sudoku solution is not valid")
	}
	return nil
}

func (s *Sudoku) checkGrid(x, y, number int8) bool {
	posX := x / 3
	posY := y / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.modified[posX+int8(i)][posY+int8(j)] == number || s.original[posX+int8(i)][posY+int8(j)] == number {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) checkVertical(y, number int8) bool {
	for _, s := range s.original {
		if s[y] == number {
			return false
		}
	}

	for _, s := range s.modified {
		if s[y] == number {
			return false
		}
	}

	return true
}

func (s *Sudoku) checkHorizontal(x, number int8) bool {
	for _, s := range s.original[x] {
		if s == number {
			return false
		}
	}

	for _, e := range s.modified[x] {
		if e == number {
			return false
		}
	}
	return true
}

func (s *Sudoku) Reset() error {
	if s == nil {
		return errors.New("can not set value to empty sudoku")
	}
	s.modified = new(SudokuField)
	return nil
}

func isNumberNotOnField(number int8) bool {
	return number < 1 || number > 9
}

func main() {
	s, err := NewSudoku(&SudokuField{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = s.Set(0, 2, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
