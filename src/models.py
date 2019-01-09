from sqlalchemy import Column, Integer, String, Unicode
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Account(Base):
    __tablename_ = 'account'
    id = Column(Integer, primary_key=True)
    email = Unicode(100)
    fname = Unicode(50, nullable=True)
    sname = Unicode(50, nullable=True)
    phone = Unicode(16, nullable=True)
    phone_code = Unicode(3, nullable=True)
    country = Unicode(50, nullable=True)
    country = Unicode(50, nullable=True)
