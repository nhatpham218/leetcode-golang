package leetcode

import (
	"strconv"
	"strings"
)

// 3484. Design Spreadsheet
// https://leetcode.com/problems/design-spreadsheet/

type Spreadsheet struct {
	cells map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		cells: make(map[string]int),
	}
}

func (s *Spreadsheet) SetCell(cell string, value int) {
	s.cells[cell] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
	s.cells[cell] = 0
}

func (s *Spreadsheet) GetValue(formula string) int {
	// Remove the '=' prefix and split by '+'
	expr := formula[1:] // Remove '='
	parts := strings.Split(expr, "+")

	// Get values for both operands
	val1 := s.parseOperand(parts[0])
	val2 := s.parseOperand(parts[1])

	return val1 + val2
}

// parseOperand returns the value of an operand (either cell reference or integer)
func (s *Spreadsheet) parseOperand(operand string) int {
	// Check if operand is a number
	if num, err := strconv.Atoi(operand); err == nil {
		return num
	}

	// Otherwise, it's a cell reference
	// If cell doesn't exist in map, return 0 (default value)
	return s.cells[operand]
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
