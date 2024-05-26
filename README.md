### Languages: [Português 🇧🇷](#api-de-pedidos-orders) | [English 🇨🇦](#orders-api)

---

# API de Pedidos (orders)

Esta API foi desenvolvida como parte do desafio do curso de Pós-Graduação em Engenharia de
Software [GoExpert](https://goexpert.fullcycle.com.br/pos-goexpert/).

- Ela contempla uma API REST, um serviço gRPC e uma interface GraphQL.
- Com duas rotas, uma para **criar um pedido** e outra para **listar todos os pedidos**.

## ⚙️ Configuração

1. Você precisará das seguintes tecnologias abaixo:
    - [Docker](https://docs.docker.com/get-docker/) 🐳
    - [Docker Compose](https://docs.docker.com/compose/install/) 🐳
    - [Postman ☄️](https://www.postman.com/downloads/) ou [VS Code](https://code.visualstudio.com/download) com a
      extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) instalada.
    - [GIT](https://git-scm.com/downloads)

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

4. A **API REST** estará disponível em `http://localhost:8000` 🚀.

5. O serviço **gRPC** estará disponível na porta `grpc://localhost:50051` 🚀.

6. A interface **GraphQL** pode ser acessada em `http://localhost:8080/graphql` 🚀.

7. Para visualizar as mensagens no **RabbitMQ**, acesse o endereço `http://localhost:15672` com as credenciais `guest`
   e `guest`.

8. Para visualizar logs do serviço `orders_api` em tempo real:
    ```sh
    docker compose logs -f orders_api
    ```

9. Para logs de todos os containers execute o comando abaixo:
   ```sh
   docker compose logs -f
   ```

10. Para limpar todos os containers, imagens e volumes do Docker, execute o comando abaixo:
    ```sh
    docker compose down -v --rmi all
    ```

## 🧪 Testes

**Utilize o arquivo `orders_api.http` como base para fazer requisições de teste.**

1. Abra o arquivo `orders_api.http` no seu editor de texto, se encontra no caminho `api/orders_api.http`.

   <br>
2. Envie requisições de teste para a API. Por exemplo, usando o VS Code, você pode instalar a
   extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).
    - Se preferir utilizar o `cURL`, você pode copiar o conteúdo do arquivo `orders_api.http` e colar
      no [Postman](https://www.postman.com/downloads/) ou terminal.
    - No VS Code, clique no botão `Send Request` que aparece ao lado de cada requisição.
    - O arquivo `orders_api.http` **contém exemplos de requisições para a API REST** para os métodos `POST` e `GET`.
   
    <br>
3. Para testar o gRPC você pode utilizar o arquivo `orders.proto` que se encontra no
   diretório `internal/infra/grpc/proto/orders.proto`,
   veja [como fazer a request pelo Postman](https://learning.postman.com/docs/sending-requests/grpc/grpc-request-interface/).
    - Uma vez importado o arquivo `orders.proto` no Postman, você pode fazer a request para o serviço **gRPC** usando o
      endpoint `grpc://localhost:50051`.
    - **Para fazer a request de criação de um pedido, utilize o método** `CreateOrder`.
      - No corpo da requisição, insira os dados do pedido no formato JSON.
      - Você pode usar o json exemplo do arquivo `orders_api.http` ou o exemplo abaixo:
        ```json
        {
          "name": "TV 49",
          "price": 4950.99,
          "tax": 1.5
        }
        ```
    - Para fazer a request de listagem de pedidos, utilize o método `ListOrders`.
   
    <br>
4. Para testar o **GraphQL** você pode utilizar a interface do **GraphQL** Playground que está disponível
   em `http://localhost:8080/graphql`. 
   - Utilize a mutation abaixo para criar um pedido:
       ```graphql
       mutation createOrder {
            createOrder(order: {name: "T-Shirt", price: 49.99, tax: 0.5}) {
                name
                price
                tax
                finalPrice
            }
       }
       ```
   - Para listar todos os pedidos, utilize a query abaixo:
       ```graphql
       query queryOrders {
            orders {
                id
                name
                price
                tax
                finalPrice
            }
       }
       ```
---

# Orders API

This API was developed as part of the challenge of the Postgraduate course in Software
Engineering [GoExpert](https://goexpert.fullcycle.com.br/pos-goexpert/).

- It includes a REST API, a gRPC service, and a GraphQL interface.
- With two routes, one to create an order and another to list all orders.

## ⚙️ Setup

1. You will need the following technologies below:
    - [Docker](https://docs.docker.com/get-docker/) 🐳
    - [Docker Compose](https://docs.docker.com/compose/install/) 🐳
    - [Postman ☄️](https://www.postman.com/downloads/) or [VS Code](https://code.visualstudio.com/download) with
      the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension installed.
    - [GIT](https://git-scm.com/downloads)

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

4. The **REST API** will be available at `http://localhost:8000` 🚀.

5. The **gRPC** service will be available on port `50051` 🚀.

6. The **GraphQL** interface can be accessed at `http://localhost:8080/graphql` 🚀.

7. To view messages in **RabbitMQ**, access `http://localhost:15672` with credentials `guest` and `guest`.

8. To view logs for the `orders_api` service in real-time:
    ```sh
    docker compose logs -f orders_api
    ```

9. To view logs for all containers:
   ```sh
   docker compose logs -f
   ```

10. To clean up all Docker containers, images, and volumes:
    ```sh
    docker compose down -v --rmi all
    ```

## 🧪 Testing

Use the `orders_api.http` file to make test requests.

1. Open the `orders_api.http` file in your text editor, located at `api/orders_api.http`.
2. Send test requests to the API. For example, using VS Code, you can install
   the [REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
    - If you prefer to use `cURL`, you can copy the contents of the `orders_api.http` file and paste it
      into [Postman](https://www.postman.com/downloads/) or terminal.
    - In VS Code, click the `Send Request` button that appears next to each request.
    - The `orders_api.http` file **contains examples of requests to the REST API** for the `POST` and `GET` methods.
    
   <br>
3. To test gRPC, you can use the `orders.proto` file located in the `internal/infra/grpc/proto/orders.proto` directory,
   see [how to make the request using Postman](https://learning.postman.com/docs/sending-requests/grpc/grpc-request-interface/).
    - Once the `orders.proto` file is imported into Postman, you can make the request to the gRPC service using
      the `grpc://localhost:50051` endpoint.
        - To make the request to create an order, use the `CreateOrder` method.
        - In the request body, enter the order data in JSON format. 
        - You can use the example JSON from the `orders_api.http` file or the example below:
          ```json
          {
            "name": "TV 49",
            "price": 4950.99,
            "tax": 1.5
          }
          ```
    - To make the request to list orders, use the `ListOrders` method.
    
    <br>
4. To test GraphQL, you can use the GraphQL Playground interface available at `http://localhost:8080/graphql`.
    - Use the mutation below to create an order:
       ```graphql
       mutation createOrder {
            createOrder(order: {name: "T-Shirt", price: 49.99, tax: 0.5}) {
                name
                price
                tax
                finalPrice
            }
       }
       ```
    - To list all orders, use the query below:
        ```graphql
        query queryOrders {
             orders {
                 id
                 name
                 price
                 tax
                 finalPrice
             }
        }
        ```