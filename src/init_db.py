import os
if os.environ.get("DEBUG", None) not in {1, "True", "true"}:
    import warnings
    warnings.filterwarnings("ignore", category=UserWarning)

import time

import models
from sqlalchemy import create_engine
from sqlalchemy.exc import OperationalError
from sqlalchemy_utils import database_exists, create_database, drop_database


def do_init():
    print("Starting DB init.")

    # Check DB is up
    engine = create_engine('postgresql://postgres:password@localhost:5432')
    while True:
        try:
            engine.connect()
            break
        except OperationalError:
            time.sleep(0.5)

    engine = create_engine('postgresql://postgres:password@localhost:5432/pairer')
    if database_exists(engine.url):
        print("Already initialized.")
        exit()

    create_database(engine.url)
    models.Base.metadata.create_all(engine)
    print("Init DB OK.")

if __name__ == '__main__':
    do_init()
