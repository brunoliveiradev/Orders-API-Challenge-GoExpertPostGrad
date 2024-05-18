**Language: [Português 🇧🇷](#api-de-pedidos-orders) | [English 🇨🇦](#orders-api)**

# API de Pedidos (orders)

## Configuração / Setup

1. Você precisará ter o Docker e Docker Compose instalados na sua máquina. Se não tiver, você pode instalar a partir dos links abaixo:
   - [Docker](https://docs.docker.com/get-docker/) 🐳
   - [Docker Compose](https://docs.docker.com/compose/install/) 🐳

2. Clone o repositório e entre no diretório do projeto.
   ```sh
   git clone https://github.com/brunoliveiradev/GoExpertPostGrad-Orders-Challenge.git
   cd GoExpertPostGrad-Orders-Challenge
   ```
   
3. Execute o comando abaixo para iniciar o ambiente de desenvolvimento:
   ```sh
   docker compose up --build -d
   ```

   Para parar os serviços:
   ```sh
   docker compose down
   ```

4. A API REST estará disponível em `http://localhost:8000/orders` 🚀.

5. O serviço gRPC estará disponível na porta `50051` 🚀.

6. A interface GraphQL pode ser acessada em `http://localhost:8080/graphql` 🚀.

7. Para visualizar logs do serviço `orders_api` em tempo real:
    ```sh
    docker compose logs -f orders_api
    ```

8. Para logs de todos os containers execute o comando abaixo:
   ```sh
   docker compose logs -f
   ```

9. Para limpar todos os containers, imagens e volumes do Docker, execute o comando abaixo:
   ```sh
   docker compose down -v --rmi all
   ```

## Testes

Utilize o arquivo `api.http` para fazer requisições de teste para a API REST e GraphQL.

1. Abra o arquivo `api.http` no seu editor de texto, se encontra no caminho `api/api.http`.
2. Envie requisições de teste para a API.


---

# Orders API

## Setup

1. You will need Docker and Docker Compose installed on your machine. If you don't have them, you can install them from the links below:
   - [Docker](https://docs.docker.com/get-docker/) 🐳
   - [Docker Compose](https://docs.docker.com/compose/install/) 🐳

2. Clone the repository and navigate to the project directory.
   ```sh
   git clone https://github.com/brunoliveiradev/GoExpertPostGrad-Orders-Challenge.git
   cd GoExpertPostGrad-Orders-Challenge
   ```
   
3. Use the Makefile to build and start the services. Run the command below to start the development environment:
   ```sh
   docker compose up --build -d
   ```

   To stop the services:
   ```sh
   docker compose down
   ```

4. The REST API will be available at `http://localhost:8000/orders` 🚀.

5. The gRPC service will be available on port `50051` 🚀.

6. The GraphQL interface can be accessed at `http://localhost:8080/graphql` 🚀.

7. To view logs for the `orders_api` service in real-time:
    ```sh
    docker compose logs -f orders_api
    ```

8. To view logs for all containers:
   ```sh
   docker compose logs -f
   ```

9. To clean up all Docker containers, images, and volumes:
   ```sh
   docker compose down -v --rmi all
   ```

## Testing

Use the `api.http` file to make test requests to the REST and GraphQL APIs.

1. Open the `api.http` file at `api/api.http` in your text editor.
2. Send test requests to the API.