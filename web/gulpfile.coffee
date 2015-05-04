'use strict';

gulp         = require 'gulp'
gutil        = require 'gulp-util'
path         = require 'path'
plumber      = require 'gulp-plumber'
browserify   = require 'browserify'
source       = require 'vinyl-source-stream'
buffer       = require 'vinyl-buffer'
less         = require 'gulp-less'
sync         = require 'browser-sync'
cssmin       = require 'gulp-cssmin'
autoprefixer = require 'gulp-autoprefixer'
sourcemaps   = require 'gulp-sourcemaps'
uglify       = require 'gulp-uglify'
argv         = require('yargs').argv


gulp.task "default", ["css", "js", "images"]

gulp.task "watch", ["default"], () ->
  sync  {
    server: { baseDir: "/opt/www" }
    files: [ "/opt/www/**" ]
    ui: false # doesn't work when we're accessing through forwarded ports
  }
  gulp.watch "/opt/src/coffee/**/*",        ['js']
  gulp.watch "/opt/src/less/**/*.less",     ['css']
  gulp.watch "/opt/src/images/**/*.png",    ['images']


gulp.task "js", () ->
  b = browserify
    entries: ["/opt/src/coffee/app.coffee"]
    extensions: [".coffee", ".cjsx"]
    debug: !argv.production
    transform: ["coffee-reactify"]

  b.bundle()
    .pipe source("app.js")
    .pipe buffer()
    .pipe sourcemaps.init { loadMaps: true }
      .pipe uglify()
      .on 'error', gutil.log
    .pipe sourcemaps.write "./"
    .pipe gulp.dest "/opt/www/js"


gulp.task "css", () ->
  gulp.src '/opt/src/less/*.less'
    .pipe plumber()
    .pipe less()
    .pipe cssmin(keepSpecialComments: 0)
    .pipe autoprefixer('last 1 version')
    .pipe gulp.dest '/opt/www/css'


gulp.task "images", () ->
  gulp.src('/opt/src/images/**/*')
    .pipe(gulp.dest('/opt/www/images'));
