
FROM python:2.7-slim

WORKDIR /app

ADD . /app

RUN pip install --trusted-host pypi.python.org -r requirements.txt

EXPOSE 5050

ENV NAME World

ENV http_proxy 127.0.0.1:6379

CMD ["python", "app.py"]
