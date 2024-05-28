package main 

import "testing"

func TestAbs(t *testing.T) {
    got := 1
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}


func TestNewtest2(t *testing.T) {
    got := 2
    if got != 2*1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

