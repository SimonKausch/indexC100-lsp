vim.lsp.start({
	name = "indexC100-lsp",
	cmd = { "./bin/indexC100-lsp" },
	root_dir = vim.fn.getcwd(),
})
