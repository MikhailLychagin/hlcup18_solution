from sqlalchemy import Column, Integer, String, Unicode, TIMESTAMP, SmallInteger, Table, ForeignKey
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship

Base = declarative_base()


class Account(Base):
    __tablename__ = 'tbl_accounts'
    id = Column(Integer, primary_key=True)
    email = Column(Unicode(100), nullable=False)
    email_domain = Column(Unicode(20), nullable=False)
    fname = Column(Unicode(50), nullable=True)
    sname = Column(Unicode(50), nullable=True)
    phone = Column(Unicode(16), nullable=True)
    phone_code = Column(Unicode(3), nullable=True)
    sex = Column(Unicode(1), nullable=False)
    birth = Column(Integer())
    birth_year = Column(SmallInteger())
    country = Column(Unicode(50), nullable=True)
    city = Column(Unicode(50), nullable=True)
    joined = Column(Integer())
    status = Column(SmallInteger())
    premium_start = Column(Integer(), nullable=True)
    premium_end = Column(Integer(), nullable=True)
    # interests = relationship("Interest", secondary="accounts_interests")
    # likes = relationship("Like", foreign_keys=["account_id"])


class Interest(Base):
    __tablename__ = 'tbl_interests'
    id = Column(Integer, primary_key=True)
    descr = Column(Unicode(100), nullable=False, unique=True)


class Like(Base):
    __tablename__ = 'tbl_likes'
    id = Column(Integer, primary_key=True)
    account_id = Column(Integer)
    liked_account_id = Column(Integer)
    ts = Column(Integer())


accounts_interests = Table(
    'accounts_interests', Base.metadata,
    Column('account_id', Integer),
    Column('interest_id', Integer),
)
