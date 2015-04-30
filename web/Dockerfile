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

RUN npm install gulp -g

RUN mkdir -p /opt/build
RUN mkdir -p /opt/www
RUN mkdir -p /opt/src

# do this first so that our cache isn't invalidated when other files change
ADD package.json /opt/build/package.json
RUN cd /opt/build && npm install
ADD gulpfile.js /opt/build/gulpfile.js
ADD gulpfile.coffee /opt/build/gulpfile.coffee

ADD src /opt/src
RUN ln -s /opt/src/html/index.html /opt/www/index.html

WORKDIR /opt/build

RUN gulp --production

ADD nginx.conf /etc/nginx/sites-available/default

VOLUME /opt/src

# Define default command.
CMD ["nginx"]

# Expose ports.
EXPOSE 80   
EXPOSE 3000 
EXPOSE 3001 

