import os
import requests

def gerar_embedding(texto):
    api_key = os.getenv("GEMINI_API_KEY")

    url = f"https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-001:embedContent?key={api_key}"

    pacote = {
        "model": "models/gemini-embedding-001",
        "content": {"parts": [{"text": texto}]},
    }

    resposta = requests.post(url, json=pacote)
    dados = resposta.json()


    vetor = dados["embedding"]["values"]
    return vetor

meu_vetor = gerar_embedding("Frete gratis em caso de defeito")

print("Quantidade de numeros gerados:", len(meu_vetor))
print("Os 5 primeiros números:", meu_vetor[:5])


