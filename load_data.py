import json
import zipfile
import datetime
import os
import re
import requests
from src.app import init_app

DATA_PATH = os.environ.get("DATA_PATH", '/tmp/data/data.zip')


def main():
    start = datetime.datetime.now()
    app = init_app()

    with zipfile.ZipFile(DATA_PATH, 'r') as zip_ref:
        for filename in zip_ref.namelist():
            if re.match(r'accounts_\d+.json$', filename):
                content_str = zip_ref.read(filename).decode('utf-8')
                content_data = json.loads(content_str).get('accounts', [])

                load_accounts_data(content_data, app)

    finish = datetime.datetime.now()
    delta = finish - start
    print(delta.total_seconds())


def load_accounts_data(content_data, app):
    s = requests.session()
    # for obj in content_data:
    #     s.post("localhost:28080/accounts/{}/".format(obj["id"]), json=obj)


if __name__ == '__main__':
    main()
