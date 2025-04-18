package handlers

import (
	"fmt"

	"github.com/SimonKausch/indexC100-lsp/mappers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(
	context *glsp.Context,
	params *protocol.CompletionParams,
) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	for word, entry := range mappers.C100Mapper {
		term := entry.Term
		description := entry.Description
		detail := fmt.Sprintf("%n\n%s", term, description)
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word,
			Detail:     &detail,
			InsertText: &term,
		})
	}
	return completionItems, nil
}
