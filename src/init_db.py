import time

import models
from sqlalchemy import create_engine
from sqlalchemy.exc import OperationalError
from sqlalchemy_utils import database_exists, create_database, drop_database


def do_init():
    engine = create_engine('postgresql://postgres:password@localhost:5432/pairer')
    while True:
        try:
            create_database(engine.url)
            break
        except OperationalError:
            time.sleep(0.5)

    models.Base.metadata.create_all(engine)
    print("Init DB OK.")

if __name__ == '__main__':
    do_init()
