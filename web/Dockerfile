FROM google/nodejs

# update the operating system
RUN apt-get update -qq && apt-get install -y build-essential
RUN apt-get install -y ruby
RUN gem install sass

RUN apt-get update -qq && apt-get -y install nginx
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

# forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log
RUN ln -sf /dev/stderr /var/log/nginx/error.log

RUN mkdir /web

RUN npm install gulp -g

# create and build all the codes
WORKDIR /web

# do this first so that our cache isn't invalidated when other files change
ADD package.json /web/package.json
RUN npm install

ADD . /web

RUN gulp

ADD nginx.conf /etc/nginx/sites-available/default

# Define default command.
CMD ["nginx"]

# Expose ports.
EXPOSE 80

