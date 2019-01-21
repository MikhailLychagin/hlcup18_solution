from sqlalchemy import Column, Integer, Unicode, SmallInteger, Table, Boolean, ARRAY, VARCHAR, UniqueConstraint
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship

Base = declarative_base()


class Account(Base):
    __tablename__ = 'tbl_accounts'
    id = Column(Integer, primary_key=True)
    email = Column(Unicode(100), nullable=False, unique=True)
    fname = Column(Unicode(50), nullable=True)
    sname = Column(Unicode(50), nullable=True)
    phone = Column(Unicode(16), nullable=True, unique=True)
    phone_code = Column(Unicode(3), nullable=True)
    sex = Column(Unicode(1), nullable=False)
    birth = Column(Integer())
    birth_year = Column(SmallInteger())
    country = Column(Unicode(50), nullable=True)
    city = Column(Unicode(50), nullable=True)
    joined = Column(Integer())
    status = Column(SmallInteger())
    premium_start = Column(Integer(), nullable=True)
    premium_finish = Column(Integer(), nullable=True)
    has_premium = Column(Boolean(create_constraint=False), nullable=True)
    interests = Column(ARRAY(VARCHAR(100)))
    __table_args__ = ({'prefixes': ['UNLOGGED']}, )


class Like(Base):
    __tablename__ = 'tbl_likes'
    id = Column(Integer, primary_key=True)
    account_id = Column(Integer)
    liked_account_id = Column(Integer)
    avg_ts = Column(Integer())
    likes_count = Column(SmallInteger())
    __table_args__ = (
        UniqueConstraint('account_id', 'liked_account_id', name="constr_likes_composed_key_unique"),
        {'prefixes': ['UNLOGGED']},
    )
