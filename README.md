# Task Manager CLI

Sample solution for the https://roadmap.sh/projects/task-tracker 

URL do projeto: https://github.com/jadersonmarc/task-manager-cli

Um gerenciador de tarefas simples em linha de comando escrito em Go.

## Visão geral

`task-manager-cli` permite adicionar, listar, atualizar o status e excluir tarefas usando um arquivo JSON local (`tasks.json`) como armazenamento.

O projeto é dividido em três camadas:

- `main.go`: entrada do aplicativo e parsing dos comandos do CLI.
- `service/`: implementa a lógica de negócio das operações de tarefas.
- `storage/`: carrega e salva tarefas em disco usando JSON.
- `task/`: define o modelo de dados da tarefa.

## Comandos

Use o binário `task-cli` seguido do comando.

- `add <descrição>`: adiciona uma nova tarefa com status `todo`
- `done <id>`: marca uma tarefa como `done`
- `progress <id>`: marca uma tarefa como `in-progress`
- `delete <id>`: remove uma tarefa
- `list`: lista todas as tarefas
- `list:done`: lista apenas tarefas concluídas
- `list:todo`: lista apenas tarefas pendentes
- `list:progress`: lista tarefas em andamento

## Exemplo de uso

```bash
# build do projeto
go build -o task-cli .

# adicionar uma tarefa
./task-cli add "Comprar leite"

# listar todas as tarefas
./task-cli list

# marcar tarefa como em progresso
./task-cli progress 1

# marcar tarefa como concluída
./task-cli done 1

# deletar tarefa
./task-cli delete 1
```

## Armazenamento

As tarefas são persistidas no arquivo `tasks.json` no diretório de execução. O armazenamento usa JSON indentado para facilitar leitura manual.

Cada tarefa possui os campos:

- `id`: identificador numérico
- `description`: descrição da tarefa
- `status`: `todo`, `in-progress` ou `done`
- `createdAt`: data de criação
- `updatedAt`: data da última atualização

## Estrutura do projeto

- `go.mod` - definição do módulo Go
- `main.go` - camada de interface do CLI
- `service/task_service.go` - lógica de adição, atualização, exclusão e listagem
- `storage/storage.go` - persistência de tarefas em `tasks.json`
- `task/task.go` - definição do tipo `Task` e status possíveis

## Requisitos

- Go 1.26+

## Observações

- O aplicativo trata IDs como inteiros e exibe erros para IDs inválidos ou tarefas não encontradas.
- Se `tasks.json` não existir, ele é criado automaticamente ao salvar a primeira tarefa.
