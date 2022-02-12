package ngrams

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

func Unigrams(line string) []string {
	return Ngrams(line, 1)
}

func Bigrams(line string) []string {
	return Ngrams(line, 2)
}

func Trigrams(line string) []string {
	return Ngrams(line, 3)
}
