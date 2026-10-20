import re
from collections import Counter


def word_count(text: str) -> dict[str, int]:
    """统计 ASCII 英文单词的频次，忽略大小写。"""
    words = re.findall(r"[A-Za-z]+", text)
    return dict(Counter(word.lower() for word in words))
