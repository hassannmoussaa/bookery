requirejs.config({
    baseUrl: (typeof _StaticFilesURLPath != "undefined" ? _StaticFilesURLPath : '/static') + '/dist/scripts/',
    waitSeconds: 0,
    paths: {
        jquery: 'lib/jquery/dist/jquery.min',
        requestjs: 'lib/requestjs/requestjs',
        dot: 'lib/dot/doT.min',
        alertjs: 'lib/alertjs/alertjs',
        tinymce: 'lib/tinymce/tinymce.min',
        cp: 'app/modules/cp',
        helpers: 'app/modules/helpers',
        urlManager: 'app/modules/urlManager',
        loader: 'app/modules/loader',
        text: 'app/modules/text',
        imgbox: 'app/modules/img-box',
        common: 'app/modules/common',
        popupjs: 'lib/popup.js/popupjs',
        moment: "lib/moment/min/moment.min",
        domurl: "lib/domurl/url.min",
        dropdown: "app/modules/dropdown",
        owl: "lib/owl.carousel/dist/owl.carousel.min",
        select2: "lib/select2/dist/js/select2.min"
    },
    urlArgs: function (id, url) {
        var args = 'app_version=' + _AppVersion;
        return (url.indexOf('?') === -1 ? '?' : '&') + args;
    }
});

require(['app']);