package tree_sitter_apex_test

import (
	"testing"

	tree_sitter_apex "github.com/rody/tree-sitter-sfapex/apex/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_apex.Language())
	if language == nil {
		t.Errorf("Error loading Apex grammar")
	}
}
