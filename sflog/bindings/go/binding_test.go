package tree_sitter_sflog_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_sflog "github.com/aheber/tree-sitter-sfapex.git/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_sflog.Language())
	if language == nil {
		t.Errorf("Error loading Sflog grammar")
	}
}
