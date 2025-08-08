# Sistema de Fila - Backend

Este projeto é um backend para um sistema de fila, desenvolvido em Go (Golang).

## Estrutura Inicial

- `cmd/` - Ponto de entrada da aplicação (main.go)
- `internal/` - Código interno do domínio
  - `handlers/` - Handlers HTTP
  - `models/` - Modelos de dados
  - `config/` - Configuração da aplicação
- `go.mod` - Gerenciamento de dependências

## Como iniciar

1. Instale o Go (https://golang.org/dl/)
2. Execute `go run ./cmd/main.go` para rodar o servidor.

## Próximos passos
- Implementar endpoints básicos de fila
- Adicionar testes
