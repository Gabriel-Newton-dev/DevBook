# DEV-BOOK É UMA REDE SOCIAL

### Possíbilitando que os usuários possam criar publicações, porém que contenha apenas texto. 

## o qual criei usando Golang, tanto Front-end quando o back-end

![Golang](img/golang2.jpeg)

### Usuários 

* CRUD
* Seguir outro usuário
* Parar de Seguir outro usuário
* Buscar todos os usuários que segue
* Buscar todos os usuários que são seguidos
* Atualuzar senha

### Publicações 

* CRUD
* Buscar publicações de acordo com os usuários que segue
* Seguir

## Nessa aplicação temos 2 componentes

* API (Back-End) 
* Web-App (Front-End)

## Banco de Dados

### Temos 2 tabelas no Database

* Usuários 
* Seguidores

## Pacotes da Aplicação

### Eu dividi em 2 tipos :

#### 1 - Pacotes Pricipais

* Main - Pacote Principal.
* Router - Irá configurar o router e todas as rotas abaixo dele.
* Controllers - Vão ficar todas as funções que iram lhe dar com as requisições HTTP, o pacote router será responsável por configurar as rotas, só que o pacote controllers que vai ter a função que vão ser executadas por essas rotas, isolei em pacotes diferentes, ela vai se comunicar ambém com pacote Modelos e Repositórios. 
* Modelos - Aonde irei guardar as 2 entidades (os Structs e metódos de Usuários e Publicações).
* Repositótios - Irá fazer interação com banco de Dados. 

#### 2 - Pacotes Auxiliares

* Config - Para configuração de variáveis de ambiente.
* Banco - para abrir a conexão com banco de dados.
* Autenticação - Cuidar do Login e criação de Token. 
* Middleware - Será usado em conjunto com pacote de autenticação, para ver se o usuário já está autenticado, requisação e a respostas. 
* Segurança - Para verificar se a senha colocada do usuário está correta com a hash do banco de dados. 
* Respostas - Para padronizar todas as respostas que a nossa API irá devolver. 

