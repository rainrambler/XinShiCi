#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
全宋词文本分割工具（修订版·支持词牌后空行接副标题）
- 按作者拆分文本，文件名格式："序号 作者.txt"
- 忽略“又”、“其一”～“其十”等常见小标题
- 词牌后紧随的非空行（可间隔若干空行）若汉字数 ≤3，视为副标题，不报警
- 输出疑似词牌/作者时附带行号
"""

import re
import sys
from collections import defaultdict

def count_chinese(text: str) -> int:
    return len(re.findall(r'[\u4e00-\u9fff]', text))

def load_set(filename: str) -> set:
    try:
        with open(filename, 'r', encoding='utf-8') as f:
            return {line.rstrip() for line in f if line.rstrip()}
    except FileNotFoundError:
        print(f"错误：文件 {filename} 不存在。", file=sys.stderr)
        sys.exit(1)

def segment_poetry(input_file: str, author_file: str, cipai_file: str) -> None:
    authors = load_set(author_file)
    cipai = load_set(cipai_file)

    # 可忽略的小标题
    exempt_titles = {'又', '其一', '其二', '其三', '其四', '其五',
                     '其六', '其七', '其八', '其九', '其十', '一', '二', '三', '四', '五',
                     '六', '七', '八', '九', '十', '十一', '十二', '十三'}

    author_order = {}
    author_contents = defaultdict(list)
    order_counter = 0
    current_author = None
    state = 'expect_author'       # 状态: expect_author | expect_author_intro | expect_cipai | expect_content
    just_had_cipai = False        # 刚处理完词牌，允许后续空行后接副标题

    with open(input_file, 'r', encoding='utf-8') as f:
        for line_no, raw_line in enumerate(f, start=1):
            line = raw_line.rstrip()

            # ---------- 空行处理 ----------
            if not line:
                if current_author is not None:
                    author_contents[current_author].append('')
                # 保留 just_had_cipai 标记，支持跨空行副标题
                continue

            # ---------- 状态机 ----------
            if state == 'expect_author':
                # 此时行为作者名
                current_author = line
                if current_author not in author_order:
                    order_counter += 1
                    author_order[current_author] = order_counter
                author_contents[current_author].append(line)
                state = 'expect_author_intro'

            elif state == 'expect_author_intro':
                # 此时行为作者介绍（第一行非空）
                author_contents[current_author].append(line)
                state = 'expect_cipai'
                # 作者介绍无论多短都不报警

            elif state == 'expect_cipai':
                # 此时行为词牌
                author_contents[current_author].append(line)
                state = 'expect_content'
                just_had_cipai = True

            elif state == 'expect_content':
                if line in authors:          # 新作者
                    current_author = line
                    if current_author not in author_order:
                        order_counter += 1
                        author_order[current_author] = order_counter
                    author_contents[current_author].append(line)
                    state = 'expect_author_intro'
                    just_had_cipai = False
                elif line in cipai or line in exempt_titles:          # 同一作者的新词牌
                    author_contents[current_author].append(line)
                    state = 'expect_cipai'   # 重新进入词牌状态，后续可接副标题
                    just_had_cipai = True
                else:
                    # 内容行（可能是副标题、正文）
                    han_cnt = count_chinese(line)

                    # 疑似短行报警（排除已知词牌、作者、豁免标题、副标题）
                    if (han_cnt <= 3
                            and line not in cipai
                            and line not in authors
                            and line not in exempt_titles
                            and not just_had_cipai):   # 词牌后的副标题不报警
                        print(f"行 {line_no}: 疑似词牌或作者: {line}")

                    author_contents[current_author].append(line)

                    # 刚处理完词牌后的第一个非空行（无论是否为副标题）后，释放标记
                    if just_had_cipai:
                        just_had_cipai = False
                    # 状态保持 expect_content

    # 写出每位作者的文件
    for author, lines in author_contents.items():
        seq = author_order[author]
        safe_author = author.replace('/', '_').replace('\\', '_')
        filename = f"{seq} {safe_author}.txt"
        with open(filename, 'w', encoding='utf-8') as out_f:
            out_f.write('\n'.join(lines) + '\n')
        print(f"已生成: {filename}")

if __name__ == '__main__':
    input_file = sys.argv[1] if len(sys.argv) > 1 else '全宋词.txt'
    segment_poetry(input_file, 'author.txt', 'cipai.txt')