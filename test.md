---
prefix: 这是一个前缀
description: 这是一个描述
---

```
snippet, err := parseSnippetFromMD(mdContent)
if err != nil {
    fmt.Fprintf(os.Stderr, "解析错误: %v\n", err)
    os.Exit(1)
}
```

