# obsidian-vscode-snippet
使用obsidian来管理vscode的代码片段



1、vscode代码片段组成部分


2、读取vscode的代码片段

/Users/th/Library/Application Support/Code/User/snippets/obsidian.code-snippets


{
    "pd": {
        "prefix": "pdhello",
        "scope": "python",
        "description": "",
        "body": [
            "import pandas as pd ",
            ""
        ]
    },
}


3、配置文件说明

```
{
    "obsidian-dir":"",  // obsidian知识库的路径
    "output-filepath": []  // 需要产出的代码片段路径，支持多个路径
}
```

4、可以使用crontab 来进行定时同步，如果着急也可以直接手动执行进行同步