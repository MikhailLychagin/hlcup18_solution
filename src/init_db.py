import models
from sqlalchemy import create_engine
from sqlalchemy_utils import database_exists, create_database, drop_database


def do_init():
    engine = create_engine('postgresql://postgres:password@localhost:5432/pairer')
    if database_exists(engine.url):
        drop_database(engine.url)
    create_database(engine.url)
    models.Base.metadata.create_all(engine)

if __name__ == '__main__':
    do_init()
