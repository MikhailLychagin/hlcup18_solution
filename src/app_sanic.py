import datetime
import io
import os
import time

import asyncpg
from asyncpg import UniqueViolationError, NotNullViolationError, InvalidTextRepresentationError
from sanic import Sanic, response

STATUS_MAP = {"свободны": 0, "заняты": 1, "всё сложно": 2}

def init_app():
    app = Sanic()
    app.config.RESPONSE_TIMEOUT = 2
    app.config.KEEP_ALIVE_TIMEOUT = 10
    app.debug = os.environ.get("DEBUG") in {"1", 1, True, "True"}
    if app.debug:
        app.date = 1545613207.0
    else:
        app.config.ACCESS_LOG = False
        app.config.DEBUG = False
        with open("/tmp/data/options.txt") as f:
            app.date = float(f.readline(1))

    return app

app = init_app()

@app.listener('before_server_start')
async def init_db(sanic, loop):
    sanic.db = await asyncpg.create_pool(host='localhost',
                                            port=5432,
                                            user='postgres',
                                            password='password',
                                            database='pairer',
                                            ssl=False,
                                            max_size=100,
                                   loop=loop)


@app.route('/accounts/new/', methods=("POST", ))
async def acc_new(request):
    global app

    body = request.json
    columns, values = io.StringIO(), io.StringIO()
    interests, likes = None, None
    i = 1
    try:
        for k, v in body.items():
            if k == "id":
                continue
            elif k == "interests":
                if v:
                    if columns.tell():
                        columns.write(",")
                    columns.write("interests")

                    if values.tell():
                        values.write(",")
                    values.write("'{")
                    values.write(
                        ",".join('"{}"'.format(interest) for interest in v)
                    )
                    values.write("}'")

                i += 1
                continue
            elif k == "likes":
                likes = v
                continue
            elif k == "status":
                v = STATUS_MAP[v]
            elif k == "birth":
                if columns.tell():
                    columns.write(",")
                columns.write("birth_year")

                parsed_date = datetime.datetime.utcfromtimestamp(v)

                if values.tell():
                    values.write(",")
                values.write(str(parsed_date.year))
            elif k == "premium":
                if columns.tell():
                    columns.write(",")
                columns.write("premium_start,premium_finish,has_premium")

                if values.tell():
                    values.write(",")
                values.write(str(v["start"]))
                values.write(",")
                values.write(str(v["finish"]))
                values.write(",")
                values.write("'t'" if v["start"] <= app.date <= v["finish"] else "'f'")

                i += 1
                continue

            if columns.tell():
                columns.write(",")
            columns.write(k)

            if values.tell():
                values.write(",")
            if isinstance(v, str):
                values.write("'")
                values.write(v)
                values.write("'")
            else:
                values.write(str(v))

            i += 1
    except (TypeError, ValueError, KeyError):
        return response.text("", status=400)

    columns.seek(0)
    values.seek(0)

    async with app.db.acquire() as conn:
        async with conn.transaction():
            try:
                account_id = await conn.fetchval(
                    """INSERT INTO tbl_accounts
                    ({})
                    VALUES ({})
                    RETURNING id;
                    """.format(columns.read(), values.read())
                )
            except (UniqueViolationError, NotNullViolationError, InvalidTextRepresentationError):
                return response.text("", status=400)

            if likes:
                await conn.executemany(
                    """INSERT INTO tbl_likes
                    (account_id, liked_account_id, avg_ts, likes_count)
                    VALUES ($1, $2, $3, 1)
                    ON CONFLICT ON CONSTRAINT constr_likes_composed_key_unique DO UPDATE
                        SET avg_ts = (cast(1 as bigint) * tbl_likes.avg_ts * tbl_likes.likes_count + $3) / (tbl_likes.likes_count + 1),
                            likes_count = tbl_likes.likes_count + 1;
                    """,
                    ((account_id, i["id"], i["ts"]) for i in likes)
                )

    return response.text("", status=201)


@app.route('/accounts/<account_id:int>/', methods=("POST", ))
async def acc_post(request, account_id):
    global app

    conn = await app.db.acquire()
    if not await conn.fetchval("""select exists(select 1 from tbl_accounts where id = $1)""", account_id):
        return response.text("", status=404)
    body = request.json
    columns, values = io.StringIO(), io.StringIO()
    likes = None
    i = 1
    try:
        for k, v in body.items():
            if k == "id":
                return response.text("", status=400)
            elif k == "status":
                v = STATUS_MAP[v]
            elif k == "likes":
                likes = v
                continue
            elif k == "interests":
                if columns.tell():
                    columns.write(",")
                columns.write("interests")

                if values.tell():
                    values.write(",")

                if v:
                    values.write("'{")
                    values.write(
                        ",".join('"{}"'.format(interest) for interest in v)
                    )
                    values.write("}'")
                else:
                    values.write("NULL")

                i += 1
                continue
            elif k == "birth":
                if columns.tell():
                    columns.write(",")
                columns.write("birth_year")

                parsed_date = datetime.datetime.utcfromtimestamp(v)

                if values.tell():
                    values.write(",")
                values.write(str(parsed_date.year))
            elif k == "premium":
                if columns.tell():
                    columns.write(",")
                columns.write("premium_start,premium_finish,has_premium")

                if values.tell():
                    values.write(",")
                values.write(str(v["start"]))
                values.write(",")
                values.write(str(v["finish"]))
                values.write(",")
                values.write("'t'" if v["start"] <= app.date <= v["finish"] else "'f'")

                i += 1
                continue
            # TODO: check "но в теле запроса переданы неизвестные поля или типы значений неверны, то ожидается код 400"

            if columns.tell():
                columns.write(",")
            columns.write(k)

            if values.tell():
                values.write(",")
            if isinstance(v, str):
                values.write("'")
                values.write(v)
                values.write("'")
            else:
                values.write(str(v))

            i += 1
    except (TypeError, ValueError, KeyError):
        return response.text("", status=400)

    columns.seek(0)
    values.seek(0)

    async with conn.transaction():
        await conn.fetchval(
            """UPDATE tbl_accounts SET
            ({}) = ({})
            WHERE id = {};
            """.format(columns.read(), values.read(), account_id)
        )

        if likes is not None:
            await conn.execute("""
            DELETE FROM tbl_likes
            WHERE account_id = $1
            """, account_id)
            if likes:
                await conn.executemany(
                    """INSERT INTO tbl_likes
                    (account_id, liked_account_id, avg_ts, likes_count)
                    VALUES ($1, $2, $3, 1)
                    ON CONFLICT ON CONSTRAINT constr_likes_composed_key_unique DO UPDATE
                        SET avg_ts = (cast(1 as bigint) * tbl_likes.avg_ts * tbl_likes.likes_count + $3) / (tbl_likes.likes_count + 1),
                            likes_count = tbl_likes.likes_count + 1;
                    """,
                    ((account_id, i["id"], i["ts"]) for i in likes)
                )

    return response.text("", status=202)


@app.route('/accounts/likes/', methods=("POST", ))
async def likes_add(request):
    global app

    try:
        body = request.json
        async with app.db.acquire() as conn:
            await conn.executemany(
                """INSERT INTO tbl_likes
                (account_id, liked_account_id, avg_ts, likes_count)
                VALUES ($1, $2, $3, 1)
                ON CONFLICT ON CONSTRAINT constr_likes_composed_key_unique DO UPDATE
                    SET avg_ts = (cast(1 as bigint) * tbl_likes.avg_ts * tbl_likes.likes_count + $3) / (tbl_likes.likes_count + 1),
                        likes_count = tbl_likes.likes_count + 1;
                """,
                ((i["liker"], i["likee"], i["ts"]) for i in body["likes"])
            )
    except (KeyError, TypeError, asyncpg.exceptions.DataError):
        return response.text("", status=400)

    return response.text("", status=202)

@app.route('/health')
async def health(request):
    return response.text("", status=200)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80, workers=4)
