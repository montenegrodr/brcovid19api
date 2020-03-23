import os
import re
import redis
import requests

from html.parser import HTMLParser

SOURCE = 'https://www.worldometers.info/coronavirus/country/brazil/'
PATTERN = '{}:\s+(.*)\s+'
CONFIRMED_ANCHOR = 'Coronavirus Cases'
DEATH_ANCHOR = 'Deaths'
RECOVERED_ANCHOR = 'Recovered'

REDIS_HOSTS = os.getenv('REDIS_HOST', 'localhost')
REDIS_PORT = os.getenv('REDIS_PORT', '6379')


class MLStripper(HTMLParser):
    def error(self, message):
        pass

    def __init__(self):
        self.reset()
        self.strict = False
        self.convert_charrefs= True
        self.fed = []

    def handle_data(self, d):
        self.fed.append(d)

    def get_data(self):
        return ''.join(self.fed)


def strip_tags(html):
    s = MLStripper()
    s.feed(html)
    return s.get_data()


def process(s):
    return int(s.strip().replace(',', ''))


def extract(anchor, text):
    retval = None
    m = re.search(PATTERN.format(anchor), text)
    if m:
        retval = process(m.group(1))
    return retval


def main():
    r = redis.Redis(host=REDIS_HOSTS, port=REDIS_PORT)
    print('Redis Address: {}:{}'.format(REDIS_HOSTS, REDIS_PORT))
    resp = requests.get(SOURCE)

    if resp.status_code == 200:
        stripped_text = strip_tags(resp.text)

        confirmed = extract(CONFIRMED_ANCHOR, stripped_text)
        deaths = extract(DEATH_ANCHOR, stripped_text)
        recovered = extract(RECOVERED_ANCHOR, stripped_text)

        if confirmed and deaths and recovered:
            r.set('data', '{};{};{}'.format(confirmed, deaths, recovered))


if __name__ == '__main__':
    main()
