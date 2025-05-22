import threading
from queue import Queue

from layers.jaccard_similarity import jaccard_similarity
from layers.ngram_overlap import ngram_overlap
from layers.normalized_levenshtein import normalized_levenshtein
from layers.sequence_matcher import sequence_matcher
from layers.tfidf_cosine import tfidf_cosine
from layers.ai_fraud_score import ai_fraud_score

def run_algorithm(func, args, results_queue, name):
    try:
        score = func(*args)
        results_queue.put((name, score))
    except Exception as e:
        print(f"[{name}] Error: {e}")
        results_queue.put((name, 0.0))

def compute_fraud_score_pipeline(question, resp1, resp2):
    algorithms = {
        "jaccard": jaccard_similarity,
        "ngram": ngram_overlap,
        "levenshtein": normalized_levenshtein,
        "sequence": sequence_matcher,
        "tfidf": tfidf_cosine,
        "ai_fraud": lambda a, b: ai_fraud_score(question, a, b)
    }

    results_queue = Queue()
    threads = []

    for name, func in algorithms.items():
        t = threading.Thread(target=run_algorithm, args=(func, (resp1, resp2), results_queue, name))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    # Collect scores
    results = {}
    sum_score = 0.0
    while not results_queue.empty():
        name, score = results_queue.get()
        results[name] = score
        sum_score += score

    return results, sum_score
