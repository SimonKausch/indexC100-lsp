package handlers

import (
	"github.com/SimonKausch/indexC100-lsp/mappers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"

	_ "github.com/tliron/commonlog/simple"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	for word, entry := range mappers.C100Mapper {
		termCopy := entry
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word,
			Detail:     &termCopy,
			InsertText: &termCopy,
		})
	}
	return completionItems, nil
}
