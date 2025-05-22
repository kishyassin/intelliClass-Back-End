from Levenshtein import distance as levenshtein

def normalized_levenshtein(text1, text2):
    lev = levenshtein(text1, text2)
    max_len = max(len(text1), len(text2))
    return 1 - (lev / max_len)