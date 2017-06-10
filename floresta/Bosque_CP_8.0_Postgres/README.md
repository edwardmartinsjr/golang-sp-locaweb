Bosque - Postgres
=================

Como carregar o bosque no postgres uando os .sql "originais"
----------------------------------

1. Compile o Dockerfile: ```docker build -t postgres:floresta .```
2. Inicie o container: ```docker run --name pg-floresta -e POSTGRES_PASSWORD=123mudar -d postgres:floresta```
3. Crie o banco de dados: ```docker exec pg-floresta createdb -U postgres -E ISO_8859_1 -T template0 --locale=pt_BR.ISO-8859-1 floresta```
4. Copie os arquivos SQL para dentro do container: ```docker cp ./*.sql pg-floresta:/tmp/```
5. Crie as tabelas: ```docker exec pg-floresta psql floresta -U postgres -f /tmp/create.sql```
6. Insira os dados: ```docker exec pg-floresta psql floresta -U postgres -f /tmp/bosque.sql```
7. Tenha paciência enquando os dados são inseridos
8. Have fun!

Como carregar o bosque usando o nosso dump
----------------------------------

1. Compile o Dockerfile: ```docker build -t postgres:floresta .```
2. Inicie o container: ```docker run --name pg-floresta -e POSTGRES_PASSWORD=123mudar -d postgres:floresta```
3. Crie o banco de dados: ```docker exec pg-floresta createdb -U postgres -E ISO_8859_1 -T template0 --locale=pt_BR.ISO-8859-1 floresta```
4. Insira o nosso dump: ```docker exec pg-floresta psql floresta -U postgres -f /tmp/bosque_cp8_postgres_dump.sql```


Observações
-----------
* O arquivo SQL original foi baixado desse [link](http://www.linguateca.pt/Floresta/ficheiros/gz/Bosque_CP_8.0.sql.gz)
