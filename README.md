# Golang Drinks 🍹
<img src="go-drink.png" alt="Golang" width="200" />

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go"/>
<img src="https://img.shields.io/badge/Excelize-Spreadsheet-blue?style=for-the-badge" alt="Excelize"/>

Pequeno utilitário em **Go** que consome uma **API pública de drinks** e gera uma **planilha Excel** com os resultados.

É um projetinho pra demonstrar:

- Consumo de API REST (HTTP + JSON) em Go  
- Modelagem de structs pra mapear respostas JSON  
- Geração de arquivo **.xlsx** usando a lib `excelize`  

---

## 🧃 O que esse projeto faz?

1. Chama a API pública [TheCocktailDB](https://www.thecocktaildb.com/) usando a primeira letra dos drinks:

   ```text
   https://www.thecocktaildb.com/api/json/v1/1/search.php?f=a

📥 Instalação / Modo de usar

- go mod init github.com/fabyo/go-drinks
- go get github.com/xuri/excelize/v2
- go mod tidy
