import json
import logging
import zipfile
import datetime
import time
import os
import re
import grequests

DATA_PATH = os.environ.get("DATA_PATH", '/tmp/data/data.zip')


def main():
    if os.environ.get("SKIP_DATA_LOAD", None) in {1, True, "True", "true", "1"}:
        print("Skipped data load.")
        exit()

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
        res = grequests.get("http://localhost:80/health").send()
        if res.response and res.response.status_code == 200:
            print("The main app is UP, loading data...")
            break
        print("Waiting for the main app to UP.")
        time.sleep(0.5)

    reqs = [grequests.post("http://localhost:80/accounts/new/", json=obj) for obj in content_data]
    load_start = time.time()
    res = grequests.map(
        reqs,
        size=800,
    )
    all_ok = all((resp.status_code < 299 for resp in res))
    if not all_ok:
        logging.info("Some data wasn't loaded successfully.")
    elapsed = time.time() - load_start
    print("Loaded {} in {}, {} IPS.".format(len(content_data), elapsed, len(content_data)/elapsed))


if __name__ == '__main__':
    main()
