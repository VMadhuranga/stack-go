package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	cases := []struct {
		operations []struct {
			method string
			data   any
		}
		expected int
	}{
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", map[string]any{"name": "Alice", "role": "Developer"}},
				{"push", map[string]any{"name": "Bob", "title": "CTO"}},
				{"size", nil},
			},
			expected: 2,
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", map[string]any{"name": "Charlie", "company": "TechCorp"}},
				{"push", map[string]any{"name": "Diana", "skills": "Python"}},
				{"push", map[string]any{"name": "Ethan", "role": "Manager"}},
				{"size", nil},
			},
			expected: 3,
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"size", nil},
			},
			expected: 0,
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", map[string]any{"name": "Frank", "experience": "5 years"}},
				{"push", map[string]any{"name": "Grace", "education": "MBA"}},
				{"push", map[string]any{"name": "Henry", "location": "New York"}},
				{"push", map[string]any{"name": "Ivy", "industry": "Finance"}},
				{"size", nil},
			},
			expected: 4,
		},
	}

	for _, c := range cases {
		s := stack{}

		for _, o := range c.operations {
			switch o.method {
			case "push":
				s.push(o.data)
			case "size":
				got := s.getSize()
				if got != c.expected {
					t.Errorf("getSize(): got: %v != expected: %v", got, c.expected)
				}
			}
		}
	}
}

func TestPopAndPeek(t *testing.T) {
	cases := []struct {
		operations []struct {
			method string
			data   any
		}
		expected []any
	}{
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", map[string]any{"name": "Alice", "role": "Developer"}},
				{"push", map[string]any{"name": "Bob", "role": "Designer"}},
				{"size", nil},
				{"peek", nil},
				{"pop", nil},
				{"size", nil},
			},
			expected: []any{
				nil,
				nil,
				2,
				map[string]any{"name": "Bob", "role": "Designer"},
				map[string]any{"name": "Bob", "role": "Designer"},
				1,
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", map[string]any{"name": "Charlie", "company": "TechCorp"}},
				{"push", map[string]any{"name": "David", "skills": []string{"Python", "JavaScript"}}},
				{"pop", nil},
				{"pop", nil},
				{"pop", nil},
			},
			expected: []any{
				nil,
				nil,
				map[string]any{"name": "David", "skills": []string{"Python", "JavaScript"}},
				map[string]any{"name": "Charlie", "company": "TechCorp"},
				nil,
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", map[string]any{"name": "Eve", "role": "Manager", "years": 5}},
				{"peek", nil},
				{"push", map[string]any{"name": "Frank", "role": "DevOps"}},
				{"size", nil},
				{"pop", nil},
				{"pop", nil},
				{"pop", nil},
			},
			expected: []any{
				nil,
				map[string]any{"name": "Eve", "role": "Manager", "years": 5},
				nil,
				2,
				map[string]any{"name": "Frank", "role": "DevOps"},
				map[string]any{"name": "Eve", "role": "Manager", "years": 5},
				nil,
			},
		},
	}

	for _, c := range cases {
		s := stack{}

		for i, o := range c.operations {
			switch o.method {
			case "push":
				s.push(o.data)
			case "pop":
				got, expected := s.pop(), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("pop(): got: %v != expected: %v", got, expected)
				}
			case "peek":
				got, expected := s.peek(), c.expected[i]
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("peek(): got: %v != expected: %v", got, expected)
				}
			case "size":
				got, expected := s.getSize(), c.expected[i]
				if got != expected {
					t.Errorf("getSize(): got: %v != expected: %v", got, expected)
				}
			}
		}
	}
}
