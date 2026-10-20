use std::collections::HashMap;

/// 统计 ASCII 英文单词的频次，忽略大小写。
pub fn word_count(text: &str) -> HashMap<String, usize> {
    let mut counts = HashMap::new();

    for word in text.split(|c: char| !c.is_ascii_alphabetic()) {
        if word.is_empty() {
            continue;
        }

        let word = word.to_ascii_lowercase();
        *counts.entry(word).or_insert(0) += 1;
    }

    counts
}

#[cfg(test)]
mod tests {
    use super::word_count;
    use std::collections::HashMap;

    #[test]
    fn test_word_count() {
        let cases: &[(&str, &str, &[(&str, usize)])] = &[
            ("empty string", "", &[]),
            ("only separators", " ,.!? \t\n", &[]),
            ("single word", "word", &[("word", 1)]),
            (
                "multiple distinct words",
                "one of each",
                &[("one", 1), ("of", 1), ("each", 1)],
            ),
            (
                "repeated words",
                "one fish two fish red fish blue fish",
                &[("one", 1), ("fish", 4), ("two", 1), ("red", 1), ("blue", 1)],
            ),
            ("case insensitive", "Go GO go gO", &[("go", 4)]),
            (
                "punctuation and whitespace",
                "car: carpet as java : javascript!!&@$%^&",
                &[
                    ("car", 1),
                    ("carpet", 1),
                    ("as", 1),
                    ("java", 1),
                    ("javascript", 1),
                ],
            ),
            (
                "consecutive separators",
                "  hello,,,   world\n\n\thello  ",
                &[("hello", 2), ("world", 1)],
            ),
            (
                "digits are separators",
                "abc123def 456",
                &[("abc", 1), ("def", 1)],
            ),
            ("apostrophe splits word", "don't", &[("don", 1), ("t", 1)]),
            (
                "non-ASCII letters are separators",
                "café 你好 naïve",
                &[("caf", 1), ("na", 1), ("ve", 1)],
            ),
        ];

        for &(name, text, expected) in cases {
            let expected: HashMap<String, usize> = expected
                .iter()
                .map(|&(word, count)| (word.to_owned(), count))
                .collect();

            assert_eq!(word_count(text), expected, "case: {name}");
        }
    }
}
