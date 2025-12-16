---
prefix: test_home
description: 这是一个描述
tag: tag1,tag2
---

```
#!/bin/bash
# 打印近三天的起止时间

for ((i=2; i>=0; i--))
do
    start_date=$(date -d "$i days ago" +"%Y%m%d")
    end_date=$(date -d "$i days ago + 1 day" +"%Y%m%d")
    echo "$start_date,$end_date"
done
```
