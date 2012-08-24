package ystocks

import (
    "testing"
)

// Ensure the Stock type can be created successfully.
func TestStockType(t *testing.T) {
    s := Stock{"INTC"}
    if s.id != "INTC" {
        t.Error("Couldn't create Stock object")
    }
}

// Ensure a single property can be grabbed from the API.
func TestGetProperty(t *testing.T) {
    s := Stock{"CSCO"}
    sym, err := s.getProperty(Symbol)

    if err != nil || sym != "CSCO" {
        t.Error("Couldn't get Symbol property")
    }
}

// Ensure multiple properties can be grabbed from the API.
func TestGetProperties(t *testing.T) {
    s := Stock{"MSFT"}
    _, err := s.getProperties([]string{Symbol, Volume})

    if err != nil {
        t.Error("Couldn't get Symbol or Volume properties")
    }

    // TODO: test values returned from method
}
