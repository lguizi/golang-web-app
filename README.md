# **GoLang WebApp**

A simple Webapp created using GoLang as Server and HTML as Front.

**Installation**

Install gorilla/mux and go-sql-driver/mysql packages. Use the _go get_ command to install the packages from GitHub like so: 
```
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```
Create the _cadastros_ table on your mysql database:
```
DROP TABLE IF EXISTS `cadastros`;
CREATE TABLE `cadastros` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `email` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```
Edit the following lines on _main.go_ putting your database user, password and database name:
```
dbUser := "" // Usuário do banco de dados
dbPass := "" // Senha do banco de dados
dbName := "" // Nome do banco de dados
```
Run _main.go_ like so:

`go run main.go`

Now you can access the page with the browser at http://localhost:8080 or using the host ip address with port 8080.

**Tasks:**
- [x] Create the server
- [x] Serve the pages
- [x] Integrate with the database
- [x] List the database contacts
- [ ] Edit contacts from database
- [ ] Remove contacts from database
- [ ] Create new database contacts

## **GoLang WebApp _(pt-BR)_**

Um Webapp simples criado usando GoLang como Server e HTML como Front.

**Instalação**

Instale os pacote gorilla/mux e go-sql-driver/mysql. Use o comando _go get_ para instalar os pacotes pelo GitHub assim:
```
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```
Create a tabela _cadastros_ no seu banco de dados mysql:
```
DROP TABLE IF EXISTS `cadastros`;
CREATE TABLE `cadastros` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `email` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```
Edite as seguintes linhas em _main.go_ colocando usuario, senha e nome do seu banco de dados:
```
dbUser := "" // Usuário do banco de dados
dbPass := "" // Senha do banco de dados
dbName := "" // Nome do banco de dados
```
Execute _main.go_ assim:

`go run main.go`

Agora você pode acessar a página com o browser em http://localhost:8080 ou usando o ip do host com a porta 8080.

**Tarefas:**
- [x] Criar o servidor
- [x] Entregar as páginas
- [x] Integrar com o banco de dados
- [x] Listar os contatos do banco de dados
- [ ] Editar contatos do banco de dados
- [ ] Remover contatos do banco de dados
- [ ] Criar novos contatos no banco de dados
