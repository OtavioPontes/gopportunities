# Gopportunities 💼

Projeto de API feita em Go para gerenciar possíveis vagas de trabalho voltadas para os desenvolvedores de software.

Nesse Projeto, é possível listar todas as vagas (oppotunities), mostrar uma vaga específica, criar a vaga, atualizar a vaga e, por fim, deletar uma vaga de trabalho. Foi criado um Logger personalizado, bem como a documentação do código via Swagger 

## Foram utilizadas as seguintes tecnologias:
- GoLang
- Gin (Framework Web)
- Gorm (ORM feito para Go)
- Swaggo (Swagger para Go)

## Como rodar o projeto
- Baixe-o
- Na pasta que o contém, rode os seguintes comandos:
```
go mod tidy
go run main.go
```

## Documentação

A documentação do CRUD foi feita com o Swagger, para acessá-la basta acessar o endpoint **/swagger/index.html**:
```
https://localhost:8080/swagger/index.html
```
