// Include gulp
var gulp = require('gulp');
// Include plugins
var gutil = require('gulp-util');
var cleanyCss = require('gulp-clean-css');
var rename = require('gulp-rename');
var compass = require('gulp-compass');
var path = require('path');
var shell = require('gulp-shell');
var minimist = require('minimist');
var uglify = require('gulp-uglify');
var pump = require('pump');
var imagemin = require('gulp-imagemin');

var paths = {
    sass: ['../src/style/scss/**/*.scss'],
    js: ['../src/scripts/**/*.js']
};

var defaultOptions = {
    string: 'dir',
    default: { dir: "" }
};

var options = minimist(process.argv.slice(2), defaultOptions);
var cssSuffix = ".min";
if (options.dir) {
    cssSuffix = "." + options.dir + cssSuffix;
}
gulp.task('default', ['compass', 'watch']);

gulp.task('compass', function () {
    gulp.src('../src/style/scss/main.scss')
        .pipe(compass({
            project: path.join(__dirname, '../'),
            css: 'dist/style/css',
            sass: 'src/style/scss'
        }))
        .on('error', function (err) {
            gutil.log('compass', err.message);
        })
        //.pipe(gulp.dest('css'))
        .pipe(cleanyCss())
        .pipe(rename({ suffix: cssSuffix }))
        .pipe(gulp.dest(path.join(__dirname, 'dist/style/css')));
});

gulp.task('buildjs', shell.task([
    'r.js -o build.js'
]));

gulp.task('compressjs', function (cb) {
  pump([
        gulp.src(paths.js),
        uglify(),
        gulp.dest(path.join(__dirname, '../dist/scripts'))
    ],
    cb
  );
});

gulp.task('imagemin', () =>
    gulp.src('../src/img/**/*')
        .pipe(imagemin([
            imagemin.gifsicle({interlaced: true}),
            imagemin.jpegtran({progressive: true}),
            imagemin.optipng({optimizationLevel: 5})
        ]))
        .pipe(gulp.dest('../dist/img'))
);

//WATCH FOR CHANGES
gulp.task('style', function () {
    gulp.watch(paths.sass, ['compass']);
});

gulp.task('scripts', function () {
    gulp.watch(paths.js, ['compressjs']);
});