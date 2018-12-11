package escaper

import (
	"unicode/utf8"
)

func Parse(str string) []string {
	const quote = "\""
	const empty = ""
	const space = " "
	const slash = "\\"

	acc := empty
	var split []string
	inQuotes := false

	completeToken := func() {
		if acc != empty {
			split = append(split, acc)
			acc = empty
		}
	}

	addToToken := func(curr string) {
		acc = acc + curr
	}

	for i, w := 0, 0; i < len(str); {
		curr, width := runeAt(str[i:])

		if i + width < len(str) {
			next, nextWidth := runeAt(str[i+width:])
			// an escaped quotation inside a quotation block
			if inQuotes &&  curr == slash && next == quote {
				addToToken(curr + next)

				// add both escaped values, skip to next
				i = i + width + nextWidth

				continue
			}
		}

		if curr == quote {
			if inQuotes {
				completeToken()
			}

			inQuotes = !inQuotes
		} else if inQuotes {
			addToToken(curr)
		} else if curr != space && !inQuotes {
			addToToken(curr)
		} else if curr == space && acc != empty {
			completeToken()
		}

		w = width
		i += w
	}

	// any remaining text
	completeToken()

	return split
}

func runeAt(str string) (string, int) {
	r, width := utf8.DecodeRuneInString(str)
	runeStr := string(r)

	return runeStr, width
}
