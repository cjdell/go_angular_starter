var fs = require('fs'),
  path = require('path'),
  gulp = require('gulp'),
  gutil = require('gulp-util'),
  jshint = require('gulp-jshint'),
  browserify = require('gulp-browserify'),
  concat = require('gulp-concat'),
  clean = require('gulp-clean'),
  beautify = require('gulp-js-beautify'),
  sass = require('gulp-sass'),
  es6ify = require('es6ify');

var externalModules = Object.keys(JSON.parse(fs.readFileSync(path.join(__dirname, 'package.json'))).dependencies);

// The are essentially two client side applications, the public frontend and the admin panel. They share very similar tasks, hence the reason for doing this...
function createAppTasks(appName, folder) {
  gulp.task(appName + '-lint', function() {
    return gulp.src(['.' + folder + '/js/**/*.js'])
      .pipe(jshint({
        esnext: true
      }))
      .pipe(jshint.reporter('default'));
  });

  gulp.task(appName + '-beautify', function() {
    return gulp.src(['./gulpfile.js', '.' + folder + '/js/**/*.js'], {
        base: '.' + folder + '/js'
      })
      .pipe(beautify({
        indent_size: 2
      }))
      .pipe(gulp.dest('.' + folder + '/js'));
  });

  // Everything required in "node_modules" is included in the external to reduce the watch build time
  gulp.task(appName + '-browserify-external', function() {
    return gulp.src(['.' + folder + '/dist/stub.js'])
      .pipe(browserify({
          require: externalModules
        })
        .on('error', function(error) {
          console.error('\u0007', error);
        }))
      .pipe(concat('external.js'))
      .pipe(gulp.dest('.' + folder + '/dist'));
  });

  gulp.task(appName + '-browserify', function() {
    return gulp.src(['.' + folder + '/js/app.js'])
      .pipe(browserify({
          debug: true,
          transform: [es6ify],
          add: [es6ify.runtime],
          external: externalModules
        })
        .on('error', function(error) {
          console.error('\u0007', error);
        }))
      .pipe(gulp.dest('.' + folder + '/dist'));
  });

  gulp.task(appName + '-sass', function() {
    return gulp.src(['.' + folder + '/css/*.scss'])
      .pipe(sass()
        .on('error', function(error) {
          console.error('\u0007', error);
        }))
      .pipe(concat('app.css'))
      .pipe(gulp.dest('.' + folder + '/dist'));
  });

  gulp.task(appName + '-watch', [appName + '-lint', appName + '-beautify', appName + '-browserify-external', appName + '-browserify', appName + '-sass'], function() {
    gulp.watch(['.' + folder + '/js/**/*.js', '.' + folder + '/js/**/*.json'], [
      appName + '-lint',
      appName + '-browserify'
    ]);

    gulp.watch(['.' + folder + '/css/**/*.scss'], [
      appName + '-sass'
    ]);
  });
}

createAppTasks('frontend', '/assets');
createAppTasks('admin', '/admin');

gulp.task('default', ['frontend-watch', 'admin-watch']);