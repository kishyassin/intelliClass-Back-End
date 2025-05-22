import difflib


def sequence_matcher(text1, text2):
    return difflib.SequenceMatcher(None, text1, text2).ratio()
