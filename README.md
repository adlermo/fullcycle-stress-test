# 🚀 Stress Test - Go Expert Challenge

CLI desenvolvido em Go para execução de testes de carga em serviços HTTP.

## 🎯 Objetivo

Executar requisições HTTP concorrentes para uma URL informada e gerar um relatório consolidado com métricas da execução.

### Parâmetros

| Flag            | Descrição                          |
| --------------- | ---------------------------------- |
| `--url`         | URL a ser testada                  |
| `--requests`    | Quantidade total de requisições    |
| `--concurrency` | Quantidade de chamadas simultâneas |

---

## 🐳 Executando com Docker

### Build

```bash
docker build -t fullcycle-stress-test .
```

### Run

```bash
docker run --rm fullcycle-stress-test \
  --url=https://app.getmonks.com \
  --requests=100 \
  --concurrency=10
```

PowerShell:

```powershell
docker run --rm fullcycle-stress-test --url=https://app.getmonks.com --requests=100 --concurrency=10
```

---

## 💻 Executando Localmente

```bash
go run ./cmd \
  --url=https://app.getmonks.com \
  --requests=100 \
  --concurrency=10
```

Ou gerando o binário:

```bash
go build -o stress-test ./cmd
```

---

## 📊 Exemplo de Saída

```text
========== STRESS TEST REPORT ==========

Total execution time: 2.21s
Total requests: 100
HTTP 200 responses: 100
Network errors: 0

Status code distribution:
  200 -> 100
```

---

## 🔀 Estratégia de Concorrência

A aplicação utiliza um Worker Pool baseado em goroutines e channels.

* Cada worker executa requisições HTTP de forma concorrente
* A quantidade de workers é definida por `--concurrency`
* O total de requisições executadas corresponde exatamente ao valor definido em `--requests`
* Os resultados são agregados para geração do relatório final

---

## 📈 Métricas Geradas

* Tempo total de execução
* Total de requisições realizadas
* Quantidade de respostas HTTP 200
* Quantidade de erros de rede
* Distribuição dos códigos HTTP retornados
