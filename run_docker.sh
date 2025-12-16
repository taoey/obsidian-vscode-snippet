docker stop obsidian-vscode-snippet && docker rm obsidian-vscode-snippet
docker run -d \
  --name obsidian-vscode-snippet \
  -p 5001:5000 \
  -v "/Users/th/Library/Application Support/Code/User/snippets/obsidian_dev.code-snippets:/home/work/obsidian-vscode-snippet/data/snippets/obsidian.code-snippets" \
  -v "/Users/th/Library/Mobile Documents/iCloud~md~obsidian/Documents/代码片段库/snippet:/home/work/obsidian-vscode-snippet/data/obsidian_data" \
  taoey/obsidian-vscode-snippet:1.0
