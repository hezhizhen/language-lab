import pytest

from wordcount.wordcount import word_count


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        pytest.param("", {}, id="empty-string"),
        pytest.param(" ,.!? \t\n", {}, id="only-separators"),
        pytest.param("word", {"word": 1}, id="single-word"),
        pytest.param(
            "one of each",
            {"one": 1, "of": 1, "each": 1},
            id="distinct-words",
        ),
        pytest.param(
            "one fish two fish red fish blue fish",
            {"one": 1, "fish": 4, "two": 1, "red": 1, "blue": 1},
            id="repeated-words",
        ),
        pytest.param("Go GO go gO", {"go": 4}, id="case-insensitive"),
        pytest.param(
            "car: carpet as java : javascript!!&@$%^&",
            {"car": 1, "carpet": 1, "as": 1, "java": 1, "javascript": 1},
            id="punctuation-and-whitespace",
        ),
        pytest.param(
            "  hello,,,   world\n\n\thello  ",
            {"hello": 2, "world": 1},
            id="consecutive-separators",
        ),
        pytest.param(
            "abc123def 456",
            {"abc": 1, "def": 1},
            id="digits-as-separators",
        ),
        pytest.param("don't", {"don": 1, "t": 1}, id="apostrophe"),
        pytest.param(
            "café 你好 naïve",
            {"caf": 1, "na": 1, "ve": 1},
            id="non-ascii-separators",
        ),
    ],
)
def test_word_count(text: str, expected: dict[str, int]) -> None:
    assert word_count(text) == expected
