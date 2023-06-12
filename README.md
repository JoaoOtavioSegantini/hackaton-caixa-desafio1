
# Descrição

Esse projeto é a resposta para o Primeiro Desafio de Desenvolvimento - Hackathon VITEC 2023 feito pela CAIXA ECONÔMICA FEDERAL.

## Instalação

Você pode rodar o projeto localmente simplesmente executando o seguinte comando no terminal:

```zsh
  docker-compose up -d
```

Como o projeto rodará dentro do container, não foi possível habilitar o https. Por isso, para que o projeto execute normalmente, a variável HTTPS_ENABLED dentro do arquivo .env deve ter um valor diferente de enabled.
Caso queria rodar o projeto fora do container do Docker, execute o seguinte comando no terminal:

```zsh
  go mod tidy
```

Logo após o download das dependências, execute o seguinte comando para subir o projeto:

```zsh
  go run framework/cmd/server/server.go
```

Lembrando que nesse último caso é possível deixar o valor da variável HTTPS_ENABLED como enabled, para habilitar os recursos do HTTP/2.


