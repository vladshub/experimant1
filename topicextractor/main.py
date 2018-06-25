import logging
import os
import socket

import requests
import spacy
from bs4 import BeautifulSoup
from elasticsearch import Elasticsearch
from facebook_business import FacebookAdsApi
# from facebookads.adobjects.adset import AdSet
from facebook_business.adobjects.adaccount import AdAccount
# from facebook_business.adobjects.adaccountuser import AdAccountUser
from facebook_business.adobjects.adset import AdSet
from facebook_business.adobjects.targetingsearch import TargetingSearch
from google.protobuf.json_format import Parse, MessageToDict
from kafka import KafkaConsumer
from kq import Job, Worker
# Set up logging
from scrapy.http import HtmlResponse

import keywee_pb2

formatter = logging.Formatter(
    fmt='[%(asctime)s][%(levelname)s] %(message)s',
    datefmt='%Y-%m-%d %H:%M:%S'
)
stream_handler = logging.StreamHandler()
stream_handler.setFormatter(formatter)
logger = logging.getLogger('keywee.topicextractor')
logger.setLevel(logging.DEBUG)
logger.addHandler(stream_handler)

my_app_id = os.environ.get('FACEBOOK_APP_ID', None)
my_app_secret =  os.environ.get('FACEBOOK_APP_SECRET', None)
my_access_token =  os.environ.get('FACEBOOK_ACCESS_TOKEN', None)  # Your user access token
my_add_account_id =  os.environ.get('FACEBOOK_AD_ACCOUNT_ID', None)
kafka_address =  os.environ.get('KAFKA_ADDRESS', 'kafka:9092')
kafka_topic = os.environ.get('KAFKA_TOPIC', 'keywee')
kafka_group_id = os.environ.get('KAFKA_GROUP_ID', 'topicextractor')
elastic_address = os.environ.get('ELASTIC_ADDRESS', 'http://elasticsearch:9200')
elastic_index = os.environ.get('ELASTIC_INDEX', 'keywee')
FacebookAdsApi.init(my_app_id, my_app_secret, my_access_token)

def processMessage(item):
    logger.info("Processing Url %s", item.url)
    content = getContent(item)
    geo_topics = getTopicsWGeo(content)
    topics = geo_topics['topics']
    geo = geo_topics['geo']
    if getFBData(item, topics, geo):
        es = Elasticsearch(hosts=elastic_address)
        es.update(index=elastic_index, doc_type='doc', id=item.id, body={'doc': MessageToDict(item)},
              params={'refresh': 'wait_for'})
    logger.info("Finished proccessing url %s", item.url)


def getFBData(item, topics, locations):
    intrests = {}
    geo = set()
    account = AdAccount(my_add_account_id)
    for t in topics:
        params = {
            'type': 'adinterest',
            'q': t,
        }
        resp = TargetingSearch.search(params=params)
        if not len(resp):
            continue
        intr = resp[0]
        intrest = intr.export_all_data()
        intrests[intrest['id']] = {
            'id': intrest['id'],
            'name': intrest['name']
        }
        logger.debug("Processed [%s] with Facebook TargetingSearch", t)

    for t in locations:
        params = {
            'type': 'adcountry',
            'q': t,
        }
        resp = TargetingSearch.search(params=params)
        if not len(resp):
            continue
        intr = resp[0]
        intrest = intr.export_all_data()
        geo.add(intrest['country_code'])
        logger.debug("Processed [%s] with Facebook TargetingSearch", t)

    user_adcluster = list(intrests.values())
    targeting_spec = {
        'geo_locations': {
            'countries': list(geo),
        },
        'user_adclusters': user_adcluster
    }

    params = {
        'targeting_spec': targeting_spec,
    }
    account_reach_estimate = account.get_reach_estimate(params=params)
    logger.debug("Got esstimate %d from facebook", len(account_reach_estimate))
    if len(account_reach_estimate) > 0:
        item.estimate_ready = account_reach_estimate[0]['estimate_ready']
        item.audiance_size = account_reach_estimate[0]['users']
        return True
    return False


def getTopicsWGeo(content):
    topics = set()
    geo = set()
    l_parser = spacy.load('en')
    parsed_content = l_parser(content)
    ents = list(parsed_content.ents)
    for entity in ents:
        if entity.label not in [393, 388, 392, 394, 391, 390]:
            t = ' '.join(t.orth_ for t in entity).strip()
            if not t or t in ['Earth', 'US', 'CO2']:
                continue
            if entity.label_ in ["PERSON", "NORP", "ORG", "PRODUCT", "EVENT", "WORK_OF_ART"]:
                topics.add(t)
            if entity.label_ in ["GPE", "LOC"]:
                geo.add(t)
    val = {}
    logger.debug("Got topics: %d geo: %d", len(topics), len(geo))
    val['topics'] = topics
    val['geo'] = geo
    return val


def getContent(item):
    user_agent = {
        'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/58: .0.3029.110 Chrome/58.0.3029.110 Safari/537.36'}
    r = requests.get(item.url, headers=user_agent)
    response = HtmlResponse(r.url, body=r.text, encoding='utf-8')


    content = cleanContent(response.css('.body-content,.article-content,#articleBody,.entry-content'))
    title = cleanContent(response.xpath("//meta[@property='og:title']/@content"))
    introduction = cleanContent(response.xpath("//meta[@property='og:description']/@content"))
    logger.debug("Got content from url %s", item.url)
    return title + " . " + introduction + " . " + content

def cleanContent(content):
    if len(content) == 0:
        return ""
    return BeautifulSoup(content[0].extract().replace("</p><p>", "</p>&nbsp;<p>").encode("UTF-8"),
    "lxml").get_text()

def deserializer(serialized):
    assert isinstance(serialized, bytes), 'Expecting a bytes'
    msg = Parse(serialized, keywee_pb2.Item())
    args = [msg]
    kwargs = {}
    return Job(
        func=processMessage,
        args=args,
        kwargs=kwargs,
    )


if __name__ == '__main__':
    logger.info("Connecting to kafka [%s] with consumer group [%s] to topic [%s]", kafka_address, kafka_group_id, kafka_topic)
    consumer = KafkaConsumer(
        bootstrap_servers=kafka_address,
        group_id=kafka_group_id,
        enable_auto_commit=False,
        auto_offset_reset='latest'
    )
    worker = Worker(
        topic=kafka_topic,
        consumer=consumer,
        deserializer=deserializer,
    )
    worker.start()
