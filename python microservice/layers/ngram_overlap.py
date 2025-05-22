from itertools import islice


def ngram_overlap(text1, text2, n=3):
    def get_ngrams(text, n):
        tokens = text.split()
        return set(zip(*[islice(tokens, i, None) for i in range(n)]))
    return len(get_ngrams(text1, n) & get_ngrams(text2, n)) / max(1, len(get_ngrams(text1, n) | get_ngrams(text2, n)))
