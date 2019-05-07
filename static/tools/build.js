({
    baseUrl: '../dist/scripts/',
    name: 'main',
    mainConfigFile: '../dist/scripts/main.js',
    out: '../dist/scripts/app.min.js',
    preserveLicenseComments: false,
    paths: {
        requireLib: 'lib/requirejs/require'
    },
    // onBuildWrite: function (moduleName, path, contents) {
    //     if (moduleName == "main") {
    //         return contents.replace('baseUrl:', 'baseUrl: _WebHost + ');
    //     }
    //     return contents;
    // },
    include: 'requireLib'
});