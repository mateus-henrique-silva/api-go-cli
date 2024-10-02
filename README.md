Aqui está um exemplo de README para o seu projeto:

---

# Application Search Server

Este projeto é um aplicativo CLI (Command Line Interface) em Go para buscar endereços IP e servidores de nome (NS) para um host especificado na internet.

## Funcionalidades

- Buscar endereços IP associados a um domínio.
- Buscar servidores de nome (NS) de um domínio.

## Pré-requisitos

- [Go](https://golang.org/dl/) instalado (versão 1.13+).
- Conexão à internet para realizar as buscas de DNS.

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/application-search-server.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd application-search-server
   ```

3. Compile o programa:

   ```bash
   go build -o searchapp main.go
   ```

## Uso

Após a compilação, o aplicativo pode ser executado para realizar buscas de IPs e servidores de nomes.

### Buscar IPs de um domínio

Para buscar os endereços IP associados a um domínio, use o comando `ip`:

```bash
./searchapp ip --host=exemplo.com
```

Se você não especificar um host, o valor padrão será `devbook.com.br`:

```bash
./searchapp ip
```

### Buscar Servidores de Nome (NS) de um domínio

Para buscar os servidores de nome (NS) de um domínio, use o comando `Servidores`:

```bash
./searchapp Servidores --host=exemplo.com
```

Se você não especificar um host, o valor padrão será `devbook.com.br`:

```bash
./searchapp Servidores
```

## Flags

- `--host`: Especifica o domínio que você deseja buscar (opcional, padrão: `devbook.com.br`).

### Exemplos de uso

1. Buscar os IPs associados ao domínio `google.com`:

   ```bash
   ./searchapp ip --host=google.com
   ```

2. Buscar os servidores de nome (NS) para `google.com`:

   ```bash
   ./searchapp Servidores --host=google.com
   ```

## Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença

Este projeto está sob a licença MIT. Consulte o arquivo `LICENSE` para mais informações.

---

Este README dá um passo a passo claro sobre como usar o aplicativo e cobriria a maior parte das necessidades básicas de um usuário.
