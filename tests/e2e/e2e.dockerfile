FROM python:3.11-slim

WORKDIR /test
COPY . .

RUN pip install --no-cache-dir -r requirements.txt

ENTRYPOINT ["pytest"]