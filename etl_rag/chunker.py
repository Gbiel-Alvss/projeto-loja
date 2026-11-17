def fatiar_texto(texto, chunk_size, chunk_overlap):
    stride = chunk_size - chunk_overlap
    chunks = []


    for i in range(0, len(texto), stride):
        pedaco = texto[i : i + chunk_size]
        chunks.append(pedaco)
    
    return chunks


with open("politicas.txt", "r", encoding="utf-8") as f:
    conteudo = f.read()

resultado = fatiar_texto(conteudo, chunk_size=200, chunk_overlap=50)

for numero, pedaco in enumerate(resultado, start=1):
    print(f"--- Chunk {numero} ---")
    print(pedaco)