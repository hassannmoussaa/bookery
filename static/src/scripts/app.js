require(['jquery', 'requestjs', 'alertjs', 'moment', 'dot', 'urlManager', 'popupjs', 'cp', 'app/modules/moment-ar', 'imgbox', 'domurl', 'dropdown', 'loader', 'common'], function ($, requestjs, alertjs, moment, doT, urlManager, popupjs) {
    alertjs.init('#alert', "ltr");
    $.ajaxSetup({
        headers: { 'X-Csrf-Token': _CSRFToken }
    });
    doT.templateSettings = {
        evaluate: /<%([\s\S]+?)%>/g,
        interpolate: /<%=([\s\S]+?)%>/g,
        encode: /<%!([\s\S]+?)%>/g,
        use: /<%#([\s\S]+?)%>/g,
        define: /<%##\s*([\w\.$]+)\s*(\:|=)([\s\S]+?)#%>/g,
        conditional: /<%\?(\?)?\s*([\s\S]*?)\s*%>/g,
        iterate: /<%~\s*(?:%>|([\s\S]+?)\s*\:\s*([\w$]+)\s*(?:\:\s*([\w$]+))?\s*%>)/,
        varname: 'it',
        strip: true,
        append: true,
        selfcontained: false
    };
    popupjs.init('#popup');
    if (_IsCPPage) {
        require(['app/cp-pages/' + _PageName.replace("cp-", "")]);
    } else {
        require(['app/pages/' + _PageName, 'app/modules/date']);
    }
});