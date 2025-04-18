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
		term := entry.Term
		description := entry.Description
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:         word,
			Detail:        &description,
			Documentation: &description,
			InsertText:    &term,
		})
	}
	for label, snippet := range mappers.CodeSnippets {
		insertText := snippet.Body
		detail := snippet.Description
		kind := protocol.CompletionItemKindSnippet
		textFormat := protocol.InsertTextFormatSnippet

		completionItems = append(completionItems, protocol.CompletionItem{
			Label:            label,
			Detail:           &detail,
			InsertText:       &insertText,
			Kind:             &kind,
			InsertTextFormat: &textFormat,
		})
	}

	return completionItems, nil
}
