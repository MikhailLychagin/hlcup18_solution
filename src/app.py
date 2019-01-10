import asyncio
import io
from datetime import datetime

import asyncpg
from aiohttp import web


ALLOWED_QUERY_FIELDS = {
    "sex_eq",
    "email_domain", "email_lt", "email_gt",
    "status_eq", "status_neq",
    "fname_eq", "fname_any", "fname_null",
    "sname_eq", "sname_starts", "sname_null",
    "phone_code", "phone_null",
    "country_eq", "country_null",
    "city_eq", "city_any", "city_null",
    "birth_year", "birth_lt", "birth_gt",
    "interests_contains", "interests_any",
    "likes_contains",
    "premium_now", "premium_null",
}

SPECIAL_ACCOUNT_FIELDS = {
    "interests", "likes"
}

STATUS_MAP = {"свободны": 0, "заняты": 1, "всё сложно": 2}


async def accounts_filter(request):
    pass


async def accounts_get(request):
    pass


async def accounts_post(request):
    account_id = int(request.match_info['id'])
    conn = await app['pool'].acquire()
    if not await conn.fetchval("""select exists(select 1 from tbl_accounts where id = $1)""", account_id):
        return web.Response(status=404)
    body = await request.json()
    columns, values = io.StringIO(), io.StringIO()
    interests, likes = None, None
    i = 1
    for k, v in body.items():
        # DELETE_ME: Поле id никогда не содержится среди обновляемых полей
        # if k == "id":
        #     continue
        if k == "interests":
            interests = v
            continue
        # DELETE_ME: лайки обновляются отдельным урлом
        # elif k == "likes":
        #     likes = v
        #     continue
        elif k == "status":
            v = STATUS_MAP[v]
        elif k == "birth":
            if columns.tell():
                columns.write(",")
            columns.write("birth_year")

            parsed_date = datetime.utcfromtimestamp(v)

            if values.tell():
                values.write(",")
            values.write(str(parsed_date.year))
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

    columns.seek(0)
    values.seek(0)

    async with conn.transaction():
        await conn.fetchval(
            """UPDATE tbl_accounts SET
            ({}) = ({})
            WHERE id = {};
            """.format(columns.read(), values.read(), account_id)
        )

        if interests:
            res = await conn.fetch(
                """
                WITH inserted_interests AS (
                    INSERT INTO tbl_interests (descr)
                    (
                        SELECT 
                            inp.descr
                        FROM
                            UNNEST($1::tbl_interests[]) AS inp
                    )
                    ON CONFLICT ON CONSTRAINT tbl_interests_descr_key DO NOTHING
                    RETURNING id
                )
                SELECT id
                FROM inserted_interests
                UNION
                SELECT id
                FROM tbl_interests
                WHERE descr = any($2::varchar[]);
                """,
                tuple((None, i,) for i in interests),
                tuple(i for i in interests)
            )
            interests_ids = ((account_id, record["id"]) for record in res)
            await conn.execute(
                """DELETE FROM accounts_interests
                WHERE account_id = $1""",
                account_id)
            await conn.executemany(
                """INSERT INTO accounts_interests (account_id, interest_id)
                VALUES ($1, $2)
                """,
                interests_ids)

    return web.Response(status=202)


async def accounts_new(request):
    body = await request.json()
    columns, values = io.StringIO(), io.StringIO()
    interests, likes = None, None
    i = 1
    for k, v in body.items():
        if k == "id":
            continue
        elif k == "interests":
            interests = v
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

            parsed_date = datetime.utcfromtimestamp(v)

            if values.tell():
                values.write(",")
            values.write(str(parsed_date.year))

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

    columns.seek(0)
    values.seek(0)

    async with app['pool'].acquire() as conn:
        async with conn.transaction():
            account_id = await conn.fetchval(
                """INSERT INTO tbl_accounts
                ({})
                VALUES ({})
                RETURNING id;
                """.format(columns.read(), values.read())
            )

            if interests:
                res = await conn.fetch(
                    """
                    WITH inserted_interests AS (
                        INSERT INTO tbl_interests (descr)
                        (
                            SELECT 
                                inp.descr
                            FROM
                                UNNEST($1::tbl_interests[]) AS inp
                        )
                        ON CONFLICT ON CONSTRAINT tbl_interests_descr_key DO NOTHING
                        RETURNING id
                    )
                    SELECT id
                    FROM inserted_interests
                    UNION
                    SELECT id
                    FROM tbl_interests
                    WHERE descr = any($2::varchar[]);
                    """,
                    tuple((None, i,) for i in interests),
                    tuple(i for i in interests)
                )
                interests_ids = ((account_id, record["id"]) for record in res)
                await conn.executemany(
                    """INSERT INTO accounts_interests (account_id, interest_id)
                    VALUES ($1, $2)
                    """,
                    interests_ids,
                )

            if likes:
                await conn.executemany(
                    """INSERT INTO tbl_likes
                    (account_id, liked_account_id, ts)
                    VALUES ($1, $2, $3);
                    """,
                    ((account_id, i["id"], i["ts"]) for i in likes)
                )

    return web.Response(status=201)


async def accounts_likes_add(request):
    try:
        body = await request.json()
        async with app['pool'].acquire() as conn:
            await conn.executemany(
                """INSERT INTO tbl_likes
                (account_id, liked_account_id, ts)
                VALUES ($1, $2, $3);
                """,
                ((i["liker"], i["likee"], i["ts"]) for i in body["likes"])
            )
    except (KeyError, TypeError, asyncpg.exceptions.DataError):
        return web.Response(status=400)

    return web.Response(status=202)


async def init_app():
    app = web.Application()
    app['pool'] = await asyncpg.create_pool(host='localhost',
                                            port=15432,
                                            user='postgres',
                                            password='password',
                                            database='pairer',
                                            ssl=False)
    app.add_routes([web.post('/accounts/new/', accounts_new)])
    app.add_routes([web.post('/accounts/likes/', accounts_likes_add)])
    app.add_routes([web.post('/accounts/{id}/', accounts_post)])

    return app


loop = asyncio.get_event_loop()
app = loop.run_until_complete(init_app())
web.run_app(app, port=28080)
