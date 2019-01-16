FROM centos:7.6.1810

ENV LANG en_US.utf8

RUN yum -y update; yum clean all
RUN yum -y install gcc make zlib-devel bzip2 bzip2-devel readline-devel sqlite sqlite-devel \
                   openssl-devel xz xz-devel libffi-devel \
                   curl-devel expat-devel gettext-devel openssl-devel zlib-devel git-core


# Python
RUN curl -L https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | bash
RUN /root/.pyenv/bin/pyenv install 3.7.2 && /root/.pyenv/bin/pyenv global 3.7.2
ARG python_interpreter=/root/.pyenv/versions/3.7.2/bin/python


# Postgresql
RUN yum -y install https://download.postgresql.org/pub/repos/yum/11/redhat/rhel-7-x86_64/pgdg-centos11-11-2.noarch.rpm
RUN yum -y install postgresql11-server sudo epel-release; yum clean all
RUN yum -y install supervisor; yum clean all
ADD ./deploy/postgres/postgresql-setup /usr/bin/postgresql-setup
#Sudo requires a tty. fix that.
RUN sed -i 's/.*requiretty$/#Defaults requiretty/' /etc/sudoers
RUN chmod +x /usr/bin/postgresql-setup
RUN /usr/bin/postgresql-setup initdb

# App
WORKDIR mhc2018_app
COPY requirements.txt .
RUN $python_interpreter -m pip install -r requirements.txt
COPY src ./src
COPY deploy ./deploy


COPY ./deploy/postgrespostgresql.conf /var/lib/pgsql/data/postgresql.conf
ADD ./deploy/postgres/pg_hba.conf /var/lib/pgsql/data/pg_hba.conf
RUN chown -R -v postgres.postgres /var/lib/pgsql/data
RUN sudo -u postgres /usr/pgsql-11/bin/postgres -D /var/lib/pgsql/data -p 5432 &
EXPOSE 80
ENTRYPOINT supervisord -n -c ./deploy/supervisord.conf
