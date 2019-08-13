package helper

import (
	"strings"
)

// 全角转半角
func ToHalfWidth(str string, num, letter, punctuation, other bool) string {
	if strings.Trim(str, " ") == "" || (!num && !letter && !punctuation && !other) {
		return str
	}

	replaceStrings := make([]string, 10)
	if num {
		numbers := map[string]string{
			"０": "0",
			"１": "1",
			"２": "2",
			"３": "3",
			"４": "4",
			"５": "5",
			"６": "6",
			"７": "7",
			"８": "8",
			"９": "9",
		}
		for k, v := range numbers {
			replaceStrings = append(replaceStrings, k, v)
		}
	}

	if letter {
		letters := map[string]string{
			"ａ": "a",
			"ｂ": "b",
			"ｃ": "c",
			"ｄ": "d",
			"ｅ": "e",
			"ｆ": "f",
			"ｇ": "g",
			"ｈ": "h",
			"ｉ": "i",
			"ｊ": "j",
			"ｋ": "k",
			"ｌ": "l",
			"ｍ": "m",
			"ｎ": "n",
			"ｏ": "o",
			"ｐ": "p",
			"ｑ": "q",
			"ｒ": "r",
			"ｓ": "s",
			"ｔ": "t",
			"ｕ": "u",
			"ｖ": "v",
			"ｗ": "w",
			"ｘ": "x",
			"ｙ": "y",
			"ｚ": "z",
			"Ａ": "A",
			"Ｂ": "B",
			"Ｃ": "C",
			"Ｄ": "D",
			"Ｅ": "E",
			"Ｆ": "F",
			"Ｇ": "G",
			"Ｈ": "H",
			"Ｉ": "I",
			"Ｊ": "J",
			"Ｋ": "K",
			"Ｌ": "L",
			"Ｍ": "M",
			"Ｎ": "N",
			"Ｏ": "O",
			"Ｐ": "P",
			"Ｑ": "Q",
			"Ｒ": "R",
			"Ｓ": "S",
			"Ｔ": "T",
			"Ｕ": "U",
			"Ｖ": "V",
			"Ｗ": "W",
			"Ｘ": "X",
			"Ｙ": "Y",
			"Ｚ": "Z",
		}
		for k, v := range letters {
			replaceStrings = append(replaceStrings, k, v)
		}
	}

	if punctuation {
		punctuations := map[string]string{
			"（": "(",
			"）": ")",
			"〔": "[",
			"〕": "]",
			"【": "[",
			"】": "]",
			"〖": "[",
			"〗": "]",
			"“": "\"",
			"”": "\"",
			"‘": "'",
			"’": "'",
			"｛": "{",
			"｝": "}",
			"《": "<",
			"》": ">",
			"：": ":",
			"。": ".",
			"，": ",",
			"、": ".",
			"；": ",",
			"？": "?",
			"！": "!",
			"…": "-",
			"‖": "|",
			"｜": "|",
			"〃": "\"",
		}
		for k, v := range punctuations {
			replaceStrings = append(replaceStrings, k, v)
		}
	}

	if other {
		others := map[string]string{
			"＄": "$",
			"＠": "@",
			"＃": "#",
			"＾": "^",
			"＆": "&",
			"＊": "*",
			"％": "%",
			"＋": "+",
			"—": "-",
			"－": "-",
			"～": "-",
		}
		for k, v := range others {
			replaceStrings = append(replaceStrings, k, v)
		}
	}

	return strings.NewReplacer(replaceStrings...).Replace(str)
}

// 判断是否为空字符
func IsEmpty(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}
