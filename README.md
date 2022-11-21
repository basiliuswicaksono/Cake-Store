# Cake-Store

### Run app in container.

To run container

```bash
docker compose up
```

To stop container

```bash
docker compose down -v
```

### App endpoint

- GET (get all cakes) : http://localhost:4000/cakes/
- GET (get cake by id) : http://localhost:4000/cakes/:{id}
- POST (add new cake) : http://localhost:4000/cakes/

```
//body
{
    "title" : "title 2",
    "description" : "description 2",
    "rating" : 3.6,
    "image" : "image 2"
}

```

- PATCH (update cake) : http://localhost:4000/cakes/:{id}

```
//body
{
    "title" : "title update",
    "description" : "description 2",
    "rating" : 3.6,
    "image" : "image 2"
}

```

- DELETE (delete cake) : http://localhost:4000/cakes/:{id}

- [postman collection](https://documenter.getpostman.com/view/12891914/2s8YmUKeyJ)
