import requests

def ai_fraud_score(question, resp1, resp2):
    prompt = f"""Tu es un professeur expert. Voici une question d'examen :

Question : {question}

Réponse A : {resp1}
Réponse B : {resp2}

Compare les deux réponses selon ces critères :
1. Similitude de style ou de structure
2. Risques de triche (copie, paraphrase)
3. Contenu inexact ou inventé

Donne une note de fraude entre 0 (aucun soupçon) et 1 (triche probable) :
Réponds seulement avec un nombre."""

    GROQ_API_URL = "https://api.groq.com/openai/v1/chat/completions"
    GROQ_API_KEY = "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"

    headers = {
        "Authorization": f"Bearer {GROQ_API_KEY}",
        "Content-Type": "application/json"
    }

    payload = {
        "model": "llama3-8b-8192",  
        "messages": [
            {"role": "user", "content": prompt}
        ]
    }

    response = requests.post(GROQ_API_URL, headers=headers, json=payload)
    response.raise_for_status()
    result = response.json()

    # Safely extract the content
    try:
        content = result["choices"][0]["message"]["content"].strip()
        score = float(content)
    except (KeyError, IndexError, ValueError):
        score = 0.0  # Fallback if parsing fails

    return score
