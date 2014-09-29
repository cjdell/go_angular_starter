var fs = require('fs'),
  path = require('path'),
  gulp = require('gulp'),
  gutil = require('gulp-util'),
  jshint = require('gulp-jshint'),
  browserify = require('gulp-browserify'),
  concat = require('gulp-concat'),
  clean = require('gulp-clean'),
  beautify = require('gulp-js-beautify'),
  jsfmt = require('gulp-jsfmt'),
  sass = require('gulp-sass'),
  es6ify = require('es6ify');

var externalModules = Object.keys(JSON.parse(fs.readFileSync(path.join(__dirname, 'package.json'))).dependencies);

gulp.task('admin-lint', function() {
  return gulp.src(['./web/admin/js/**/*.js'])
    .pipe(jshint({
      esnext: true
    }))
    .pipe(jshint.reporter('default'));
});

gulp.task('admin-beautify', function() {
  return gulp.src(['./gulpfile.js', './web/admin/js/**/*.js'], {
      base: './web/admin/js'
    })
    .pipe(beautify({
      indent_size: 2
    }))
    .pipe(gulp.dest('./web/admin/js'));
});

gulp.task('admin-jsfmt', function() {
  return gulp.src(['./gulpfile.js', './web/admin/js/**/*.js'], {
      base: './web/admin/js'
    })
    .pipe(jsfmt.format())
    .pipe(gulp.dest('./web/admin/js'));
});

// Everything required in "node_modules" is included in the external to reduce the watch build time
gulp.task('admin-browserify-external', function() {
  return gulp.src(['./web/admin/dist/stub.js'])
    .pipe(browserify({
        require: externalModules
      })
      .on('error', function(error) {
        console.error('\u0007', error);
      }))
    .pipe(concat('external.js'))
    .pipe(gulp.dest('./web/admin/dist'));
});

gulp.task('admin-browserify', function() {
  return gulp.src(['./web/admin/js/app.js'])
    .pipe(browserify({
        debug: true,
        transform: [es6ify],
        add: [es6ify.runtime],
        external: externalModules
      })
      .on('error', function(error) {
        console.error('\u0007', error);
      }))
    .pipe(gulp.dest('./web/admin/dist'));
});

gulp.task('admin-sass', function() {
  return gulp.src(['./web/admin/css/*.scss'])
    .pipe(sass()
      .on('error', function(error) {
        console.error('\u0007', error);
      }))
    .pipe(concat('app.css'))
    .pipe(gulp.dest('./web/admin/dist'));
});

gulp.task('admin-watch', ['admin-lint', 'admin-beautify', 'admin-browserify-external'], function() {
  gulp.watch(['./web/admin/js/**/*.js', './web/admin/js/**/*.json'], [
    'admin-lint',
    'admin-browserify'
  ]);

  gulp.watch(['./web/admin/css/**/*.scss'], [
    'admin-sass'
  ]);
});