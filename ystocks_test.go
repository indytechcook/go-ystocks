package ystocks

import (
  "testing"
)

func TestStockType(t *testing.T) {
  s := Stock{"INTC"}
  if s.id != "INTC" {
    t.Error("Couldn't create Stock object")
  }
}
