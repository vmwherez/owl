This folder wants to be a repo of the ASGI app.
For now letting all `ipynb` live here.

This is what I've been doing to start it locally:

```
minikube start
uvicorn app_old:app --reload
```

### Dockerfile

```docker

FROM python:3.8-slim-buster
EXPOSE 8000

# Add demo app
COPY ./app /app
WORKDIR /app

RUN python -m pip install --upgrade pip

# Install gunicorn & falcon
RUN pip install -r requirements.txt

CMD ["uvicorn", "app:app"]

```

### Quarto

VSCode -> ... (menu) -> render as HTML

https://quarto.org/docs/get-started/
https://quarto.org/docs/get-started/hello/vscode.html

### Typora & Obsidian
