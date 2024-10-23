# Guia de Instalação do PostgreSQL no WSL

## Índice
1. [Ativando o WSL no Windows](#ativando-o-wsl-no-windows)
2. [Instalando o PostgreSQL no WSL](#instalando-o-postgresql-no-wsl)
3. [Configurando o PostgreSQL para Iniciar Automaticamente](#configurando-o-postgresql-para-iniciar-automaticamente)
4. [Configurando o PostgreSQL para Aceitar Conexões Externas](#configurando-o-postgresql-para-aceitar-conexões-externas)

---

## Ativando o WSL no Windows

1. Abra o **Windows PowerShell** como administrador.
2. Execute o seguinte comando para habilitar o WSL:

    ```powershell
    wsl --install
    ```

   Isso instalará a última versão do WSL e o Ubuntu como distribuição padrão.

3. Reinicie o computador quando solicitado.

4. Após a reinicialização, abra novamente o PowerShell e confirme se o WSL está ativo com o comando:

    ```powershell
    wsl -l -v
    ```

5. Se necessário, atualize o kernel do WSL:

    ```powershell
    wsl --update
    ```

---

## Instalando o PostgreSQL no WSL

1. Abra o terminal do WSL (Ubuntu) a partir do menu iniciar.
2. Atualize os pacotes do sistema:

    ```bash
    sudo apt update && sudo apt upgrade -y
    ```

3. Instale o PostgreSQL com o seguinte comando:

    ```bash
    sudo apt install postgresql postgresql-contrib -y
    ```

4. Verifique se a instalação foi bem-sucedida e o serviço está em execução:

    ```bash
    sudo service postgresql status
    ```

---

## Configurando o PostgreSQL para Iniciar Automaticamente

O WSL não mantém serviços em execução após ser fechado. Para garantir que o PostgreSQL inicie sempre que o WSL for iniciado, siga os seguintes passos:

1. Abra o arquivo `.bashrc` para edição:

    ```bash
    nano ~/.bashrc
    ```

2. Adicione a seguinte linha no final do arquivo para iniciar o PostgreSQL sempre que o WSL for iniciado:

    ```bash
    sudo service postgresql start
    ```

3. Salve e feche o arquivo (`Ctrl + O`, `Enter`, e `Ctrl + X`).

4. Carregue as alterações do `.bashrc` sem reiniciar o WSL:

    ```bash
    source ~/.bashrc
    ```

Agora, o PostgreSQL será iniciado automaticamente ao abrir o WSL.

---
## Criando o usuario e o Banco de dados

Para poder acessar de fora não se pode utilizar o usuario root pois o serviço de autenticação ssl nao permite a utilização se nao for um ssl emitido pela propria maquina virtual.

1. Para criar um novo usuário no PostgreSQL, você precisará acessar o terminal do WSL e utilizar o comando CREATE USER dentro do postgres. Vou te guiar pelo processo passo a passo:
    - **Commando no WSL para acessar ao Terminal Postgres**
    ```bash
        sudo -u postgres psql
    ```

    - **Comando no Terminal Postgres para crirar o Usuario e o DB**
    ```sql
    DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'admin') THEN
            CREATE USER admin WITH PASSWORD 'adminadmin' SUPERUSER;
        END IF;
        IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'ancora') THEN
            CREATE DATABASE ancora OWNER admin;
        END IF;
    END $$;
    ```
---

## Configurando o PostgreSQL para Aceitar Conexões Externas

Por padrão, o PostgreSQL está configurado para aceitar conexões apenas locais. Para permitir que ele aceite conexões da máquina Windows, siga os seguintes passos:

### Passo 1: Configurando o `postgresql.conf`

1. Abra o arquivo de configuração `postgresql.conf` para edição:

    ```bash
    sudo nano /etc/postgresql/12/main/postgresql.conf
    ```

    **Nota:** Verifique a versão do PostgreSQL instalada em `/etc/postgresql/`.

2. Encontre a linha com `listen_addresses` e descomente-a, alterando seu valor para:

    ```bash
    listen_addresses = '*'
    ```

3. Salve e feche o arquivo (`Ctrl + O`, `Enter`, e `Ctrl + X`).

### Passo 2: Configurando o `pg_hba.conf`

1. Abra o arquivo de configuração `pg_hba.conf` para editar as permissões de acesso:

    ```bash
    sudo nano /etc/postgresql/12/main/pg_hba.conf
    ```

2. No final do arquivo, adicione a seguinte linha para permitir conexões de fora da rede do WSL (use o IP do host WSL, que geralmente é `172.19.0.1`, mas pode variar):

    ```bash
    host    all             all             172.19.0.1/16           md5
    ```

3. Salve e feche o arquivo.

### Passo 3: Reiniciando o PostgreSQL

1. Após as alterações de configuração, reinicie o serviço do PostgreSQL:

    ```bash
    sudo service postgresql restart
    ```

---

## Verificando a Conexão Externa

Agora você pode tentar acessar o PostgreSQL a partir da máquina Windows utilizando um cliente de banco de dados (como DBeaver ou pgAdmin) ou pelo terminal `psql` no WSL. O endereço IP do servidor será `localhost` ou o IP atribuído ao WSL.

### Exemplo de Conexão:

1. Abra o PowerShell no Windows.
2. Utilize o comando `psql` para se conectar ao PostgreSQL no WSL:

    ```bash
    psql -h localhost -U postgres -d nome_do_banco
    ```

Se tudo estiver configurado corretamente, você deverá ser capaz de se conectar ao banco de dados PostgreSQL rodando dentro do WSL diretamente a partir do Windows.

---

## Troubleshooting

1. **Erro de conexão:** Verifique se o IP da rede WSL é correto. Utilize o comando abaixo para verificar o IP do WSL:

    ```bash
    ip addr show eth0
    ```

2. **Configuração de Firewall:** Certifique-se de que o firewall do Windows permite conexões na porta 5432 (ou a porta que você configurou para o PostgreSQL).

---
