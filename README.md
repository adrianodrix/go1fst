# Meu Projeto Go

Um servidor web simples desenvolvido em Go para demonstrar conceitos básicos de programação, pacotes e módulos.

## 🚀 Funcionalidades

- **Servidor HTTP** usando Gorilla Mux
- **API REST** com endpoints para status e informações
- **Pacote próprio** (`utils`) com funções utilitárias
- **Timestamp automático** em todas as mensagens
- **Estrutura modular** demonstrando organização de código

## 📋 Pré-requisitos

- Go 1.24 ou superior
- Git

## 🛠️ Instalação

1. Clone o repositório:
```bash
git clone https://github.com/adrianodrix/go1fst.git
cd go1fst
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Execute o servidor:
```bash
go run main.go
```

## 🌐 Endpoints

### GET /
Retorna uma mensagem de boas-vindas com timestamp.

**Resposta:**
```
[2025-07-02 11:02:42] Olá! Este é meu primeiro servidor web em Go!
```

### GET /api/status
Retorna o status do servidor em formato JSON.

**Resposta:**
```json
{
  "status": "online",
  "timestamp": "2025-07-02 11:02:42",
  "message": "Servidor funcionando!"
}
```

### GET /api/info
Retorna informações internas do sistema.

**Resposta:**
```json
{
  "info": "função interna - chamada externamente"
}
```

## 📁 Estrutura do Projeto

```
meu-projeto/
├── main.go          # Arquivo principal do servidor
├── utils/
│   └── helpers.go   # Pacote com funções utilitárias
├── go.mod           # Definição do módulo e dependências
├── go.sum           # Checksums das dependências
└── README.md        # Este arquivo
```

## 🔧 Dependências

- **github.com/gorilla/mux** - Router HTTP para Go
- **pacote utils** - Funções utilitárias próprias

## 🎯 Conceitos Demonstrados

- **Módulos Go** - Gerenciamento de dependências
- **Pacotes próprios** - Organização de código
- **Visibilidade** - Funções públicas vs privadas
- **Imports** - Como importar pacotes em módulos
- **Servidor HTTP** - Criação de APIs REST

## 🚀 Como Executar

```bash
# Executar o servidor
go run main.go

# Acessar no navegador
http://localhost:8080
```

## 📝 Desenvolvimento

Para adicionar novas funcionalidades:

1. Crie novos arquivos no diretório `utils/` para funções utilitárias
2. Adicione novos endpoints no `main.go`
3. Use `go mod tidy` para gerenciar dependências

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

**Adriano Santos**
- GitHub: [@adrianodrix](https://github.com/adrianodrix)
- Email: [hello@adrianodrix.me](mailto:hello@adrianodrix.me)

## 🙏 Agradecimentos

- Comunidade Go
- Gorilla Mux por fornecer um excelente router HTTP
- Este projeto foi criado como parte do aprendizado da linguagem Go

---

⭐ Se este projeto te ajudou, considere dar uma estrela no repositório!