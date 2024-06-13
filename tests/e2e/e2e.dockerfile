FROM python:3.12-slim

ENV POETRY_NO_INTERACTION=1 \
    POETRY_VIRTUALENVS_CREATE=false \
    POETRY_CACHE_DIR='/var/cache/pypoetry' \
    POETRY_HOME='/usr/local' \
    POETRY_VERSION=1.8.2

RUN python -m pip install poetry

WORKDIR /test
COPY . .

RUN poetry install --no-interaction --no-ansi

ENTRYPOINT ["pytest"]