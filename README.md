## Go-Transfer: Uma API para Transferências Financeiras

![Go](https://img.shields.io/badge/Go-v1.20+-00ADD8?style=flat-square&logo=go)  
![Docker](https://img.shields.io/badge/Docker-latest-2496ED?style=flat-square&logo=docker)  
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-latest-336791?style=flat-square&logo=postgresql)

---

### 👀 Visão Geral

Go-Transfer é uma API RESTful construída em Go que permite a criação e execução de transferências financeiras entre usuários.  
O projeto adota uma arquitetura modular, separando as responsabilidades em camadas bem definidas (API, Configuração, Domínio e Infraestrutura) para promover a manutenibilidade, testabilidade e escalabilidade.

---

### ⚙️ Funcionalidades Principais

- Criação de Usuários
- Transferências Financeiras com verificação de saldo e consistência transacional
- Notificações via serviço HTTP externo (simulado)
- Arquitetura orientada a domínio (DDD simplificado)

---

### 🧱 Estrutura do Projeto

- `cmd/` → Ponto de entrada da aplicação (`main.go`)
- `internal/`
    - `api/` → Handlers HTTP
    - `config/` → Setup de dependências
    - `domain/`
        - `entities/` → `User`, `Wallet`, `Transaction`, `Notification`
        - `port/` → Interfaces do domínio
        - `usecase/` → Regras de negócio
    - `env/` → Variáveis de ambiente
    - `infra/`
        - `database/` → GORM + PostgreSQL
        - `externals/` → Integração com APIs externas (autorização, notificação)
        - `repositories/` → Implementações concretas dos repositórios

---
### 🛠️ Princípios e Padrões de Projeto

- Dependency Injection (Utilizado nos Controllers e Services via construtor, promovendo Inversão de Controle)
- Repository Pattern (Aplicado na camada de acesso a dados — ex: UserRepository, TransactionRepository, etc.)
- Service Layer Pattern (ou Use Case Layer) (Presente na camada de lógica de negócio — ex: UserService, AuthUseCase, etc.)
- Factory Pattern (Utilizado de forma implícita nas funções que constroem instâncias — ex: NewUserService(), NewQueueClient(), etc.)
- Strategy Pattern (Aplicado na implementação de mensageria — ex: diferentes estratégias de envio como SQS, Kafka, HTTP)
- YAGNI (Evita implementação de funcionalidades não necessárias — código criado apenas quando há demanda real)
- KISS (Manutenção de lógica clara e simples, sem overengineering — funções pequenas, coesas e legíveis)
- DRY (Reutilização de lógica através de helpers, serviços centralizados e abstrações comuns)
- Law of Demeter (Respeitado ao evitar que uma classe acesse internals de outras diretamente — services falam com interfaces, não com implementações internas)
- Composition over Inheritance (Uso de composição ao invés de herança — ex: injeção de dependências e uso de interfaces para comportamento dinâmico)
- SRP (Cada módulo tem uma única responsabilidade bem definida — ex: AuthService cuida apenas de autenticação)
- OCP (Código aberto para extensão, fechado para modificação — novos métodos ou implementações adicionados sem alterar os existentes)
- LSP (Interfaces são substituíveis por suas implementações — ex: NotificationSender pode ser Email, SMS ou Push)
- ISP (Interfaces pequenas e específicas — ex: Consumer, Publisher, Validator, etc.)
- DIP (Classes de alto nível não dependem de implementações, mas de abstrações — ex: services dependendo de interfaces)

### 📋 Pré-requisitos

- Go 1.20+
- PostgreSQL
- (Opcional) Docker e Docker Compose

---

### 🔧 Configuração

Crie um arquivo `.env` com o seguinte conteúdo:

```
PORT=":8080"
NOTIFICATION_BASE_URL=https://util.devi.tools/api/v1/notify
AUTHORIZATION_BASE_URL=https://util.devi.tools/api/v2/authorize

DATABASE_URL=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=go-transfer

```

Certifique-se de que o PostgreSQL esteja rodando.

---

### 🚀 Execução

**Sem Docker:**

```bash
git clone https://github.com/seu-usuario/go-transfer.git
cd go-transfer
go mod tidy
go run cmd/main.go
```

**Com Docker Compose:**

```bash
docker-compose up -d
```

---

### 📌 Endpoints

**POST /users**

```json
{
  "name": "João",
  "document": "12345678900",
  "email": "joao@email.com",
  "type": "COMMON"
}
```

**POST /transfers**

```json
{
  "sender_id": 1,
  "receiver_id": 2,
  "amount": 100.50
}
```

---

### ✅ Testes

```bash
go test ./...
```

---

### 🎯 Melhorias Futuras

- [x] Separação em camadas (API, domínio, infraestrutura)
- [x] Testes unitários com banco em memória
- [ ] Autenticação/autorização JWT
- [ ] Cache com Redis
- [ ] Swagger/OpenAPI
- [ ] Integração com mensageria (Kafka, RabbitMQ)

---
