package ngrams

// Ngrams returns N-grams for input string, e.g.:
// Ngrams("Hello!", 2) -> [He el ll lo o!]
func Ngrams(line string, n int) []string {
	if n <= 0 {
		return []string{}
	}

	if len(line) < n {
		return []string{}
	}

	ngrams := make([]string, len(line)-n+1)

	for i := 0; i < len(line)-n+1; i++ {
		ngrams[i] = line[i : i+n]
	}

	return ngrams
}

// Unigrams returns 1-grams for input string, e.g.:
// Unigrams("Hello!") -> [H e l l o !]
func Unigrams(line string) []string {
	return Ngrams(line, 1)
}

// Bigrams returns 2-grams for input string, e.g.:
// Bigrams("Hello!") -> [He el ll lo o!]
func Bigrams(line string) []string {
	return Ngrams(line, 2)
}

// Trigrams returns 3-grams for input string, e.g.:
// Trigrams("Hello!") -> [Hel ell llo lo!]
func Trigrams(line string) []string {
	return Ngrams(line, 3)
}
