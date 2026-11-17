 API de Pedidos, Agente IA (Tool Calling) & Servidor MCP

API REST em Go desenvolvida com Gin e GORM (SQLite) para gestão de pedidos e faturamento, integrada a um agente de LLM com Tool Calling e um servidor MCP (Model Context Protocol) sobre SSE/HTTP.

 Arquitetura

O projeto foi dividido em arquivos separados para cada parte cuidar de uma única tarefa:

- `main.go`: Inicia o projeto, abre o banco de dados, liga o servidor MCP em segundo plano (porta `8081`) e sobe as rotas da API com o Gin (porta `8080`).
- `config.go`: Lê a chave da API e a porta no arquivo `.env` e impede o servidor de subir se a chave estiver faltando.
- `models.go`: Define as estruturas dos dados do sistema (pedidos, itens, recibo) e os formatos de conversa usados para falar com a IA.
- `repository.go`: É o único arquivo que mexe no banco de dados (SQLite com GORM), cuidando de criar as tabelas, salvar e buscar pedidos.
- `service.go`: Cuida apenas das contas e regras da loja (validar pedido, calcular o frete de R$ 15,00 e somar o faturamento total), sem saber nada de internet ou JSON.
- `handlers.go`: Recebe as requisições HTTP do usuário, confere se os dados vieram no formato certo, chama o serviço ou o agente e devolve a resposta em JSON.
- `ai_client.go`: Faz a chamada HTTP para a API da Google/OpenAI enviando a pergunta e devolvendo a resposta.
- `agent.go`: Conecta a IA às funções da loja — envia o catálogo de ferramentas, executa o cálculo de faturamento quando a IA pede e devolve o resultado para ela responder ao usuário.
- `mcp_server.go`: Sobe um servidor MCP (via SSE/HTTP) para que outros programas de IA também consigam consultar o faturamento da loja.
Pré-requisitos

- Go 1.21+
- Chave de API do Google Gemini (`GEMINI_API_KEY`)

Como Rodar

1. Clone o repositório e acesse a pasta do projeto.
2. Crie um arquivo `.env` a partir do `.env.example` e preencha sua chave:
   env
   GEMINI_API_KEY=sua_chave_aqui
   PORT=8080
3. Baixe as dependências:
go mod tidy
4. Execute a aplicação:
go run .

 Endpoints

| Método | Rota | Porta | Descrição 

| `GET` | `/health` | `8080` | Verifica se a API está ativa (`{"status": "UP"}`). 
| `POST` | `/pedidos` | `8080` | Valida e registra um novo pedido com seus itens no banco, retornando o recibo com frete. 
| `GET` | `/pedidos` | `8080` | Lista todos os pedidos salvos com seus respectivos itens. 
| `GET` | `/pedidos?id={id}` | `8080` | Busca um pedido específico pelo ID (`400` se inválido, `404` se não encontrado). 
| `GET` | `/faturamento` | `8080` | Retorna a quantidade total de pedidos e o valor total faturado (itens + frete). 
| `POST` | `/ia/perguntar` | `8080` | Envia uma pergunta ao agente de IA, que aciona ferramentas locais quando necessário. 
| `GET` | `/sse` | `8081` | Canal SSE/HTTP do servidor MCP para conexão de clientes de IA. 


 Exemplos de Uso 

 Criar um pedido

```powershell
Invoke-RestMethod -Uri http://localhost:8080/pedidos -Method Post -ContentType "application/json" -Body '{"pedido_id": 101, "cliente": "Lucas Lima", "items": [{"nome": "Teclado", "preco": 350.0}]}'

```

 Consultar faturamento via Agente IA

```powershell
(Invoke-WebRequest -Uri http://localhost:8080/ia/perguntar -Method Post -ContentType "application/json" -Body '{"pergunta": "Qual o faturamento total da loja ate agora e quantos pedidos foram feitos?"}').Content

```
