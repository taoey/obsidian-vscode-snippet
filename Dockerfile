FROM python:3.11-slim

WORKDIR /home/work/obsidian-vscode-snippet

# 1 安装必要工具
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        supervisor \
        cron \
        ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# 2 拷贝依赖文件，并安装依赖
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


#设置定时任务
COPY mycron /etc/cron.d/mycron
RUN chmod 0644 /etc/cron.d/mycron && \
    crontab /etc/cron.d/mycron

# supervisord 配置文件
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# 3 拷贝python 代码: 【web】下的子目录到【app-py】 中
COPY web web
COPY config config
COPY data data
COPY output/obsidian-vscode-snippet_linux_amd64 /home/work/obsidian-vscode-snippet/app-exe

# 4 设置启动命令
WORKDIR /home/work/obsidian-vscode-snippet
EXPOSE 5000
CMD ["/usr/bin/supervisord", "-n"]
