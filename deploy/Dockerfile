FROM centos:7.6.1810

RUN yum -y install gcc make zlib-devel bzip2 bzip2-devel readline-devel sqlite sqlite-devel \
                   openssl-devel xz xz-devel libffi-devel \
                   curl-devel expat-devel gettext-devel openssl-devel zlib-devel git-core
RUN curl -L https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | bash
RUN /root/.pyenv/bin/pyenv install 3.7.2 && /root/.pyenv/bin/pyenv global 3.7.2