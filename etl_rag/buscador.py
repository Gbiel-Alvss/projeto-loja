import os
import requests
import math

def similaridade_cosseno(v1, v2):
    produto = sum(a * b for a,b in zip(v1,v2))
    tamanho_v1 = math.sqrt(sum(a * a for a in v1))
    tamanho_v2 = math.sqrt(sum(b * b for b in v2))
    return produto / (tamanho_v1 * tamanho_v2)

def fatiar_texto(texto, chunk_size, chunk_overlap):
    stride = chunk_size - chunk_overlap
    chunks = []

    for i in range(0, len(texto), stride):
        pedaco = texto[i: i + chunk_size]
        chunks.append(pedaco)

    return chunks

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

with open("politicas.txt", "r", encoding="utf-8") as f:
    conteudo = f.read()

pergunta = "Se a peça veio estragada de fabrica, quem paga o envio dos correios?"

vetor_pergunta = gerar_embedding(pergunta)

chunks = fatiar_texto(conteudo, 200, 50)

for numero, chunk in enumerate(chunks, start=1):
    vetor_chunk = gerar_embedding(chunk)
    nota = similaridade_cosseno(vetor_pergunta, vetor_chunk)
    print(f"Chunk {numero} | Nota: {nota:.4f} | Texto: {chunk}\n")




