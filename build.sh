#!/bin/bash

# build.sh - Go 交叉编译脚本
# 用法: ./build.sh

set -e  # 遇到错误立即退出

# === 配置区 ===
# 项目名（最终二进制文件名前缀）
APP_NAME="obsidian-vscode-snippet"

# 输出目录
DIST_DIR="output"

# 支持的平台和架构组合 (格式: GOOS/GOARCH)
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"   # Intel macOS
    "darwin/arm64"   # Apple Silicon (M1/M2)
    "windows/amd64"
    "windows/arm64"
)

# === 构建逻辑 ===
echo "🚀 开始构建 $APP_NAME 的多平台二进制文件..."

# 创建输出目录
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

# 获取当前目录作为源码路径（假设 main 包在当前目录）
SRC_PATH="."

# 遍历所有平台进行构建
for platform in "${PLATFORMS[@]}"; do
    IFS='/' read -r GOOS GOARCH <<< "$platform"

    # 设置输出文件名
    OUTPUT_NAME="$APP_NAME"
    if [[ "$GOOS" == "windows" ]]; then
        OUTPUT_NAME+=".exe"
    fi
    FINAL_NAME="${APP_NAME}_${GOOS}_${GOARCH}${OUTPUT_NAME##$APP_NAME}"

    echo "📦 构建: $GOOS/$GOARCH -> $DIST_DIR/$FINAL_NAME"

    # 执行交叉编译
    CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" go build \
        -ldflags="-s -w" \
        -o "$DIST_DIR/$FINAL_NAME" \
        "$SRC_PATH"
done

echo "✅ 构建完成！所有二进制文件位于: ./$DIST_DIR/"
ls -lh "$DIST_DIR"