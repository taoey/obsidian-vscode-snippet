from flask import Flask, render_template, request, jsonify
import json
import os
from datetime import datetime

app = Flask(__name__)
JSON_FILE = '../config/config.json'

# 初始化默认数据（如果文件不存在）
DEFAULT_DATA = {
    "obsidian_dir": "",
    "output_filepath": [],
    "no_need_convert_special_char": [
        "TM_SELECTED_TEXT",
        "TM_CURRENT_LINE",
        "TM_CURRENT_WORD",
        "TM_LINE_INDEX",
        "TM_LINE_NUMBER",
        "TM_FILENAME",
        "TM_FILENAME_BASE",
        "TM_DIRECTORY",
        "TM_FILEPATH",
        "RELATIVE_FILEPATH",
        "CLIPBOARD",
        "WORKSPACE_NAME",
        "WORKSPACE_FOLDER",
        "CURSOR_INDEX",
        "CURSOR_NUMBER",
        "CURRENT_YEAR",
        "CURRENT_YEAR_SHORT",
        "CURRENT_MONTH",
        "CURRENT_MONTH_NAME",
        "CURRENT_MONTH_NAME_SHORT",
        "CURRENT_DATE",
        "CURRENT_DAY_NAME",
        "CURRENT_DAY_NAME_SHORT",
        "CURRENT_HOUR",
        "CURRENT_MINUTE",
        "CURRENT_SECOND",
        "CURRENT_SECONDS_UNIX",
        "CURRENT_TIMEZONE_OFFSET"
    ]
}

if not os.path.exists(JSON_FILE):
    with open(JSON_FILE, 'w', encoding='utf-8') as f:
        json.dump(DEFAULT_DATA, f, ensure_ascii=False, indent=2)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/data')
def get_data():
    try:
        if os.path.exists(JSON_FILE):
            with open(JSON_FILE, 'r', encoding='utf-8') as f:
                data = json.load(f)
        else:
            data = DEFAULT_DATA.copy()
        return jsonify(data)
    except Exception as e:
        return jsonify({'error': '读取失败: ' + str(e)}), 500

@app.route('/save', methods=['POST'])
def save_data():
    try:
        # 获取前端传来的字段
        obsidian_dir = request.form.get('obsidian_dir', '').strip()
        output_text = request.form.get('output_filepath', '').strip()
        special_text = request.form.get('no_need_convert_special_char', '').strip()

        # 处理 output_filepath：按行分割，过滤空行
        output_list = [line.strip() for line in output_text.split('\n') if line.strip()]

        # 处理 no_need_convert_special_char：同上
        special_list = [line.strip() for line in special_text.split('\n') if line.strip()]

        new_data = {
            "obsidian_dir": obsidian_dir,
            "output_filepath": output_list,
            "no_need_convert_special_char": special_list
        }

        with open(JSON_FILE, 'w', encoding='utf-8') as f:
            json.dump(new_data, f, ensure_ascii=False, indent=2)

        return jsonify({'message': '✅ 配置已成功保存到服务器！'})

    except Exception as e:
        return jsonify({'error': '❌ 保存失败: ' + str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, host='127.0.0.1', port=5000)