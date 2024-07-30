package main

import (
	"testing"
)

func TestAccept(t *testing.T) {
	subjects := []string{"Math", "English", "Science"}
	grades := []int{85, 92, 78}
	expected := map[string]int{
		"Math":    85,
		"English": 92,
		"Science": 78,
	}
	result := accept(subjects, grades)
	for key, val := range expected {
		if result[key] != val {
			t.Errorf("For key %v, expected %v but got %v", key, val, result[key])
		}
	}
}

func TestAvg(t *testing.T) {
	dict := map[string]int{
		"Math":    85,
		"English": 92,
		"Science": 78,
	}
	expected := (85 + 92 + 78) / 3
	result := Avg(dict, 3)
	if result != expected {
		t.Errorf("Expected average %v but got %v", expected, result)
	}
}

func TestStuGrade(t *testing.T) {
	dict := map[string]int{
		"Math":    85,
		"English": 92,
		"Science": 78,
	}
	expected := map[string]string{
		"Math":    "B",
		"English": "A",
		"Science": "C",
	}
	result := stuGrade(dict)
	for key, val := range expected {
		if result[key] != val {
			t.Errorf("For key %v, expected grade %v but got %v", key, val, result[key])
		}
	}
}
