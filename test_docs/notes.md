# Mountain

## Data Sources

### `iypnb` rendered to HTML with Quarto

https://quarto.org/docs/get-started/hello/vscode.html

### notes or docs folder

```python
import os
import markdown
import falcon

# Define a Falcon resource for serving Markdown files as HTML
class MarkdownResource:
    def on_get(self, req, resp, filename):
        # Check that the file exists
        filepath = os.path.join('path/to/markdown/files', filename)
        if not os.path.exists(filepath):
            raise falcon.HTTPNotFound()

        # Convert the Markdown file to HTML
        with open(filepath, 'r') as f:
            content = f.read()
            html = markdown.markdown(content)

        # Set the response content type and body to the HTML
        resp.content_type = 'text/html'
        resp.body = html

# Create a Falcon app and add the Markdown resource to it
app = falcon.App()
app.add_route('/markdown/{filename}', MarkdownResource())

# Run the Falcon app on port 8000
if __name__ == '__main__':
    from wsgiref import simple_server
    httpd = simple_server.make_server('localhost', 8000, app)
    httpd.serve_forever()
```

### NOAA API

- https://www.weather.gov/documentation/services-web-api

### RSS

### Beautiful Soup

#### Set up Falcon and ASGI

- on metal
- Docker

#### Confirm 'hello world' 200

#### Serve Static HTML

#### Pandas and SQLite

- also CSV, xlsx

#### Redis Cache

#### Websockets

#### Deployment & Scale

##### On metal: Gunicorn and IPTables

#### Kubernetes

Minikube

​ - Ansible??

- postgres

- NoSQL

## ASGI: async Python

- https://asgi.readthedocs.io/en/latest/introduction.html
- https://github.com/MagicStack/uvloop
- https://github.com/python-hyper/wsproto
- https://www.infoworld.com/article/3658336/asgi-explained-the-future-of-python-web-development.html

`9:02p`

## falcon

- https://stackoverflow.com/questions/34821974/how-to-serve-a-static-webpage-from-falcon-application
- https://www.uvicorn.org/#quickstart
- https://falcon.readthedocs.io/en/stable/user/quickstart.html
- https://github.com/falconry/falcon
- https://falcon.readthedocs.io/en/stable/api/media.html
- https://stackoverflow.com/questions/34821974/how-to-serve-a-static-webpage-from-falcon-application

https://stackoverflow.com/questions/59821476/api-rest-falcon-python-post-method this would be synchronous request

moving forward with Falcon and Uvicorn... will need to come back to orchestration (ngnix k3s, etc)

```
uvicorn things_asgi:app
```

## Logging

- https://realpython.com/python-logging/
- https://www.datadoghq.com/blog/python-logging-best-practices/

## Redis

_Thought I would use it for log cache, but probably a ton of uses_

https://github.com/jedp/python-redis-log/blob/master/redislog/handlers.py

https://developer.redis.com/explore/what-is-redis/

- https://realpython.com/python-redis/
- https://docs.redis.com/latest/rs/references/client_references/client_python/

### SQLalchemy

https://realpython.com/python-sqlite-sqlalchemy/

https://www.digitalocean.com/community/tutorials/how-to-use-python-markdown-with-flask-and-sqlite

`9:58p` Successfully recalling sqlite database and serving raw data from falcon

`10:29p` `PostHandler()`

https://falcon.readthedocs.io/en/stable/api/media.html

don't forget to await media

```
async def on_post(self, req, resp):
	obj = await req.get_media()
	message = obj.get("info")
	print(message)
```

`10:52p`

##### [Python and Redis Tutorial - Caching API Responses](https://www.youtube.com/watch?v=_8lJ5lp8P0U "Python and Redis Tutorial - Caching API Responses")

https://realpython.com/python-redis/

### Scaling and deployment

In a single instance, one could use `gunicorn` in front of gunicorn.

> If you have a cluster of machines with Kubernetes, Docker Swarm Mode, Nomad, or another similar complex system to manage distributed containers on multiple machines, then you will probably want to handle replication at the cluster level instead of using a process manager (like Gunicorn with workers) in each container.

- https://github.com/berndverst/falcon-docker
- https://www.uvicorn.org/deployment/
- https://fastapi.tiangolo.com/deployment/docker/

- running server behind ngnix

## Concurrency and Parallelism in Python

- https://realpython.com/python-concurrency/
- https://realpython.com/python-gil/
- https://pyo3.rs/v0.13.2/parallelism.html
- https://stackoverflow.com/questions/60148992/running-python-code-in-parallel-from-rust-with-rust-cpython

[[notes]]

[index.html](/)

Math
[[MATH]]

$$
\int_{0}^{n}
$$

[microhooks](/microhooks)

## uvicorn

- https://www.uvicorn.org/#quickstart

### ASGI

- https://asgi.readthedocs.io/en/latest/introduction.html

## TODOS

- [ ] enumerate features
- [ ] reimplement with better commits, refactor, add tests
- [ ] there is a data upload feature, a Weather API, map, reads markdown and ipynb...
- [ ] Redis in DSIR by downey
- [ ] collect network engineering docs
- [ ] collect stats and ML docs
- [ ] question of where additional files should live, gitignore, etc
- [ ] docstrings are in the wrong place throughout
- [ ] more clarity around quarto now that nbconvert works well on-the-fly

```

```
