import asyncio

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


async def accounts_filter(request):
    for k in request.query.keys():
        if k not in ALLOWED_QUERY_FIELDS:
            return web.Response(status=400)




async def init_app():
    app = web.Application()
    app['pool'] = await asyncpg.create_pool(host='localhost',
                                            port=15432,
                                            user='postgres',
                                            password='password',
                                            ssl=False)
    app.add_routes([web.get('/accounts/filter/', accounts_filter)])

loop = asyncio.get_event_loop()
app = loop.run_until_complete(init_app())
web.run_app(app)
