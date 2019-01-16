import asyncio
import uvloop
asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
import io
import os
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


class AppHandlers:
    def __init__(self, app) -> None:
        self._app = app
        super().__init__()

    async def accounts_filter(self, request):
        pass


    async def accounts_get(self, request):
        pass


    async def accounts_post(self, request):
        account_id = int(request.match_info['id'])
        conn = await self._app['pool'].acquire()
        if not await conn.fetchval("""select exists(select 1 from tbl_accounts where id = $1)""", account_id):
            return web.Response(status=404)
        body = await request.json()
        columns, values = io.StringIO(), io.StringIO()
        likes = None
        i = 1
        try:
            for k, v in body.items():
                if k == "id":
                    return web.Response(status=400)
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

                    parsed_date = datetime.utcfromtimestamp(v)

                    if values.tell():
                        values.write(",")
                    values.write(str(parsed_date.year))
                elif k == "premium":
                    if columns.tell():
                        columns.write(",")
                    columns.write("premium_start,premium_end,has_premium")

                    if values.tell():
                        values.write(",")
                    columns.write(str(v["premium_start"]))
                    columns.write(",")
                    columns.write(str(v["premium_end"]))
                    columns.write(",")
                    columns.write("1" if v["premium_start"] <= self._app['date'] <= v["premium_end"] else "0")

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
            return web.Response(status=400)

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

        return web.Response(status=202)


    async def accounts_new(self, request):
        body = await request.json()
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

                    parsed_date = datetime.utcfromtimestamp(v)

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
                    values.write("'t'" if v["start"] <= self._app['date'] <= v["finish"] else "'f'")

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
            return web.Response(status=400)

        columns.seek(0)
        values.seek(0)

        async with self._app['pool'].acquire() as conn:
            async with conn.transaction():
                account_id = await conn.fetchval(
                    """INSERT INTO tbl_accounts
                    ({})
                    VALUES ({})
                    RETURNING id;
                    """.format(columns.read(), values.read())
                )

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

        return web.Response(status=201)


    async def accounts_likes_add(self, request):
        try:
            body = await request.json()
            async with self._app['pool'].acquire() as conn:
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
            return web.Response(status=400)

        return web.Response(status=202)


    async def health(self, request):
        return web.Response(status=200)


async def init_app():
    app = web.Application()
    app['debug'] = os.environ.get("DEBUG") in {"1", 1, True, "True"}
    if app['debug']:
        app['date'] = 1545613207.0
    else:
        with open("/tmp/data/options.txt") as f:
            app['date'] = float(f.readline(1))
    app['pool'] = await asyncpg.create_pool(host='localhost',
                                            port=5432,
                                            user='postgres',
                                            password='password',
                                            database='pairer',
                                            ssl=False,
                                            max_size=100)
    app_handlers = AppHandlers(app)
    app.add_routes([web.post('/accounts/new/', app_handlers.accounts_new)])
    app.add_routes([web.post('/accounts/likes/', app_handlers.accounts_likes_add)])
    app.add_routes([web.post('/accounts/{id}/', app_handlers.accounts_post)])
    app.add_routes([web.get('/health', app_handlers.health)])

    return app


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    app = loop.run_until_complete(init_app())
    web.run_app(app, port=80, access_log=None)
