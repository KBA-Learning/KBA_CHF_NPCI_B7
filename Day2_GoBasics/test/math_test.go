package main

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
    result := Add(3, 2)
    require.Equal(t, 6, result, "Add(3, 2) should be 5")
}

func TestSubtract(t *testing.T) {
    result := Subtract(3, 2)
    require.Equal(t, 1, result, "Subtract(3, 2) should be 1")
}