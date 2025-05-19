package tree_sitter_soql_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_soql "github.com/aheber/tree-sitter-sfapex.git/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_soql.Language())
	if language == nil {
		t.Errorf("Error loading Soql grammar")
	}
}
