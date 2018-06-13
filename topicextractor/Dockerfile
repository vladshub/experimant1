FROM python:3.6

RUN apt-get update && apt-get install libsnappy-dev -y

COPY ./requirements.txt /app/requirements.txt
WORKDIR /app
RUN pip install -r requirements.txt
RUN python -m spacy download en

COPY . /app

ENV FACEBOOK_APP_ID=${FACEBOOK_APP_ID}
ENV FACEBOOK_APP_SECRET=${FACEBOOK_APP_SECRET}
ENV FACEBOOK_ACCESS_TOKEN=${FACEBOOK_ACCESS_TOKEN}
ENV FACEBOOK_AD_ACCOUNT_ID=${FACEBOOK_AD_ACCOUNT_ID}
ENV KAFKA_ADDRESS=kafka:9092
ENV KAFKA_TOPIC=keywee
ENV KAFKA_GROUP_ID=${KAFKA_GROUP_ID}
ENV ELASTIC_ADDRESS=elasticsearch:9200
ENV ELASTIC_INDEX=keywee

CMD ["python", "main.py"]