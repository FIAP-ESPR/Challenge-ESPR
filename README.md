# Challenge-ESPR

|Nome|RM|
|:-:|:-:|
|**Diogo Julio** | **[ RM 553837 ]**|
|**Matheus Zottis** | **[ RM 94119 ]**|
|**Victor Didoff** | **[ RM 552965 ]**|
|**Vinicius Silva** | **[ RM 553240 ]**|

>  Temos uma paginas com alguns tutoriais separados por aquivo em [[ Challeng-ESPR ]](./docs/)

## Run Project

### Requisitos
- Instale o **Go 1.23.0** ou superior.
- Instale e Crie o Banco de Dados do Projeto. [[ WSL DB Tutorial ]](./docs/DataBase.md) | [[ DB command ]](./db/sql/script.sql)

### Instruções para Executar o Projeto

1. **Clone o repositório**
   ```bash
   git clone https://github.com/FIAP-ESPR/Challenge-ESPR.git
   ```

2. **Navegue até o diretório raiz do projeto**
   ```bash
   cd Challenge-ESPR
   ```

3. **Executar o projeto diretamente**
    ```bash
    go run .
    ```

    Ou, se preferir gerar um executável

4. **Compilar e executar o binário**
    ```bash
    go build .
    ```

### Troubleshooting

1. **Verifique o arquivo de configuração**
    
    Em [/conf/db.conf](./conf/db.conf) garanta que esta prenchido com o os dados do seu banco.

    Seu **[arquivo](./conf/db.conf)** deve conter essa estrutura
    ```JSON
    {
        "db" : "",
        "port" : "",
        "username" : "",
        "password" : "",
        "database" : ""
    }
    ```

2. **Verifique o arquivo do Tutorial do WSL**

    Pode ser um possivel erro de conexão com o Banco de Dados [**[ WSL DB Tutorial ]**](./docs/DataBase.md)