package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieInsert(t *testing.T) {
	root := newNode()

	root.insert("a")
	root.insert("aba")
	root.insert("abc")
	root.insert("ba")
	root.insert("bc")
	root.insert("ą")

	var expected = &node{
		children: map[rune]*node{
			'a': {
				children: map[rune]*node{
					'b': {
						children: map[rune]*node{
							'c': {
								children: make(map[rune]*node),
								eow:      true,
							},
							'a': {
								children: make(map[rune]*node),
								eow:      true,
							},
						},
						eow: false,
					},
				},
				eow: true,
			},
			'b': {
				children: map[rune]*node{
					'a': {
						children: make(map[rune]*node),
						eow:      true,
					},
					'c': {
						children: make(map[rune]*node),
						eow:      true,
					},
				},
				eow: false,
			},
			'ą': {
				children: make(map[rune]*node),
				eow:      true,
			},
		},
		eow: false,
	}

	assert.Equalf(t, expected, root, "inserted root and equivalent struct mismatch")
}

func TestTrieSearch(t *testing.T) {
	root := newNode()

	root.insert("a")
	root.insert("aba")
	root.insert("abc")
	root.insert("ba")
	root.insert("bc")
	root.insert("ą")

	assert.ElementsMatchf(t, []string{"a", "aba", "abc"}, root.search("a"), "elements for prefix 'a' mismatch")
	assert.ElementsMatchf(t, []string{"aba", "abc"}, root.search("ab"), "elements for prefix 'ab' mismatch")
	assert.ElementsMatchf(t, []string{"abc"}, root.search("abc"), "elements for prefix 'abc' mismatch")
	assert.ElementsMatchf(t, []string{"ą"}, root.search("ą"), "elements for prefix 'ą' mismatch")
	assert.ElementsMatchf(t, []string{"a", "aba", "abc", "ba", "bc", "ą"}, root.search(""), "all inserted elements mismatch")
	assert.ElementsMatchf(t, nil, root.search("abcd"), "array should have been empty")
	assert.ElementsMatchf(t, nil, root.search("c"), "array should have been empty")
}
