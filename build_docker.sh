set -xu
# 生成python依赖
pip install pipreqs
pipreqs . --force --ignore .env,.history,__pycache__,venv

./build.sh

# 构建并运行docker容器
export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890
docker build -t taoey/obsidian-vscode-snippet:1.0 .
