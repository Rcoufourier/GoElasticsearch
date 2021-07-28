# GoElasticsearch


## I - Intro

### 1 - Goal

*GoElasticsearch is a school project made for @HETIC, with the aim to create a library of books and searching feature using the GoogleBooks API and Elastic Search.*

### 2 - Stack

| Service        | Type                                              |
| ---------- | -------------------------- |
| Golang          | Back-end language                    |
| ElasticSearch | Non-relational database |
| Docker      | Virtual Container                        |



## II - Running the Project

This project uses Docker containers, so first, we have to run Docker

```
docker-composer build

docker-composer up -d
```


To make requests we use Postman, install Postman and then create a new request.
Make sure to copy this requests with the correct HTTP method, then send.

#### Add entries from the API

First of all, we have to populate our database with the data from the GoogleBooks API.


```
GET http://localhost:8080/add-from-api
```

#### Search all books from database

To look up for every books in the database:

```
GET http://localhost:8080/all-books
```

#### Search books from database using a keyword

To look up for books corresponding to a specific keyword, will return all the matching results:

```
GET http://localhost:8080/search?keyword={keyword}
```

#### Add a single book

To manually add a book to the database, proceed like this:

```
POST http://localhost:8080/add-book
```

The request's body should look like this:

```json
{
"title": "Titre du roman",
"author": "toto",
"abstract": "Ceci est un résumé"
}
```

#### Edit book

If you want to edit a book by its ID, follow these instructions:

```
PUT http://localhost:8080/edit/{id}
```

```json
{
"title": "Titre changé",
"author": "toto toto",
"abstract": "Ceci est un résumé modifié"
}
```

#### Delete book

To delete a book by ID in the database, launch a request using the **DELETE** method and the corresponding url

```
DELETE http://localhost:8080/delete/{id}
```

#### Display logs
```
docker container logs {containerName}
```

## Team
- Camille Arsac

- James Bissick

- Remi Coufourier
