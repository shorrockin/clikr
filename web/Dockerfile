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

WORKDIR /opt/build

# manually npm install stuff here. by doing this instead of using a
# package.json file we can more efficiently re-build our containers
# as the docker cache will be more fine grained, so adding or changing
# specific packages will not require a full big npm install. with that
# in mind, add new packages to the bottom when changing to limit the
# impact to the container cache
RUN npm install browser-sync@^2.6.9
RUN npm install coffee-script@^1.9.2
RUN npm install vinyl-source-stream@^1.1.0
RUN npm install gulp@^3.8.11
RUN npm install gulp-util@^3.0.4
RUN npm install gulp-autoprefixer@^2.2.0
RUN npm install gulp-cssmin@^0.1.7
RUN npm install gulp-concat@^2.5.2
RUN npm install gulp-plumber@^1.0.0
RUN npm install gulp-watch@^4.2.4
RUN npm install gulp-changed@^1.2.1
RUN npm install gulp-open@^0.3.2
RUN npm install gulp-less@^3.0.3
RUN npm install coffeeify@^1.1.0
RUN npm install browserify@^10.0.0
RUN npm install uglifyify@^3.0.1
RUN npm install yargs@^1.3.3
RUN npm install react@^0.13.2
RUN npm install flux@^2.0.3
RUN npm install jquery@^2.1.4
RUN npm install superagent@^1.2.0
RUN npm install gulp-sourcemaps@^1.5.2
RUN npm install vinyl-buffer@^1.0.0
RUN npm install gulp-uglify@^1.2.0
RUN npm install coffee-reactify@^3.0.0
RUN npm install object-assign@^2.0.0

# install the gulp files needed to run all our gulp commands
ADD gulpfile.js /opt/build/gulpfile.js
ADD gulpfile.coffee /opt/build/gulpfile.coffee

# the node modules needs to be a parent of the gulp source directory
# in order for browserify to find it, so we create a symlink in our
# main directory so it all works out
RUN ln -s /opt/build/node_modules /opt/node_modules

ADD src /opt/src
RUN ln -s /opt/src/html/index.html /opt/www/index.html

RUN gulp --production

ADD nginx.conf /etc/nginx/sites-available/default

VOLUME /opt/src

# Define default command.
CMD ["nginx"]

# exposes the nginx and browser-sync ports
EXPOSE 80   
EXPOSE 3000 
