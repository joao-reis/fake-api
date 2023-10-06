# Welcome to FakeAPI

Thank you for choosing FakeAPI for your development needs.

You can use FakeAPI to mock any api you need, FakeAPI will serve a HTTP server with the API defined in the `config.yaml` file.

## Define your API

Change the file `config.yaml` accordingly to your API configuration.

E.g.

```yaml
addr: "0.0.0.0"
port: 8000
api:
  - 
    endpoint: /
    verb: POST
    content_type: "application/json"
    response: |
      "{
        'id': '1',
        'title': 'Blue Train',
        'artist': 'John Coltrane',
        'price': 56.99
      }"
    status_code: 201
  - 
    endpoint: /albums
    verb: GET
    headers:
      - '"Accept-Encoding", "*"'
    content_type: "application/json"
    response: |
      [
          {
            'id': '1',
            'title': 'Blue Train',
            'artist': 'John Coltrane',
            'price': 56.99
          },
          {
            'id': '2',
            'title': 'Jeru',
            'artist': 'Gerry Mulligan',
            'price': 17.99
          },
          {
            'id': '3',
            'title': 'Sarah Vaughan and Clifford Brown',
            'artist': 'Sarah Vaughan',
            'price': 39.99
          }
      ]
    status_code: 200
  - 
    endpoint: /albums/:id
    verb: GET
    headers:
      # - '"Access-Control-Allow-Origin", "*"'
      # - '"Access-Control-Max-Age", "30"'
    content_type: "application/json"
    response: "{'id': '1', 'title': 'Blue Train', 'artist': 'John Coltrane', 'price': 56.99}"
    status_code: 200
```

## Starting the API Server

Execute the following command to start your API server:
```console
make run
```

## Making API calls

You can start doing HTTP calls to your API.

E.g.

```console
curl http://localhost:8000/albums
```
