'use strict';

gulp       = require 'gulp'
gp         = (require 'gulp-load-plugins') lazy: false
path       = require 'path'
browserify = require 'browserify'
source     = require 'vinyl-source-stream'
less       = require 'gulp-less'
sync       = require 'browser-sync'
argv       = require('yargs').argv


gulp.task "default", ["css", "js", "images"]

gulp.task "watch", ["default"], () ->
  sync  {
    server: { baseDir: "/opt/www" }
    files: [ "/opt/www/**" ]
    ui: false # doesn't work when we're accessing through forwarded ports
  }
  gulp.watch "/opt/src/coffee/**/*.coffee", ['js']
  gulp.watch "/opt/src/less/**/*.less",     ['css']
  gulp.watch "/opt/src/images/**/*.png",    ['images']


gulp.task "js", () ->
  browserify
    entries: ["/opt/src/coffee/app.coffee"]
    extensions: [".coffee", ".js"]
    debug: !argv.production
  .transform "coffeeify"
  .transform "deamdify"
  .transform "uglifyify"
  .bundle()
  .pipe source "app.js"
  .pipe gulp.dest "/opt/www/js"


gulp.task "css", () ->
  gulp.src '/opt/src/less/*.less'
    .pipe gp.plumber()
    .pipe less()
    .pipe gp.cssmin keepSpecialComments: 0
    .pipe gp.autoprefixer 'last 1 version'
    .pipe gulp.dest '/opt/www/css'


gulp.task "images", () ->
  gulp.src('/opt/src/images/**/*')
    .pipe(gulp.dest('/opt/www/images'));
