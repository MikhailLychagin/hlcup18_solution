import json
import zipfile
import datetime
import time
import os
import re
import requests
import grequests

DATA_PATH = os.environ.get("DATA_PATH", '/tmp/data/data.zip')


def main():
    start = datetime.datetime.now()
    print("Started data load.")

    with zipfile.ZipFile(DATA_PATH, 'r') as zip_ref:
        for filename in zip_ref.namelist():
            if re.match(r'accounts_\d+.json$', filename):
                content_str = zip_ref.read(filename).decode('utf-8')
                content_data = json.loads(content_str).get('accounts', [])

                load_accounts_data(content_data)

    finish = datetime.datetime.now()
    delta = finish - start
    print(delta.total_seconds())


def load_accounts_data(content_data):
    while True:
        try:
            res = requests.get("http://localhost:80/health")
            res.raise_for_status()
        except requests.RequestException:
            print("Waiting for the main app to UP.")
            time.sleep(0.5)
            continue

        if res.status_code == 200:
            print("The main app is UP, loading data...")
            break

    reqs = [grequests.post("http://localhost:80/accounts/new/", json=obj) for obj in content_data]
    load_start = time.time()
    res = grequests.map(
        reqs,
        size=400,
    )
    assert all((resp.status_code < 299 for resp in res))
    print("Loaded pack in {}".format(time.time() - load_start))


if __name__ == '__main__':
    main()
