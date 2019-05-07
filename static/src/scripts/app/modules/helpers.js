define([], function () {
    //DETECT MOBILE DEVICES
    var isMobile = {
        Android: function () {
            return navigator.userAgent.match(/Android/i);
        },
        BlackBerry: function () {
            return navigator.userAgent.match(/BlackBerry/i);
        },
        iOS: function () {
            return navigator.userAgent.match(/iPhone|iPad|iPod/i);
        },
        Opera: function () {
            return navigator.userAgent.match(/Opera Mini/i);
        },
        Windows: function () {
            return navigator.userAgent.match(/IEMobile/i);
        },
        any: function () {
            return (isMobile.Android() || isMobile.BlackBerry() || isMobile.iOS() || isMobile.Opera() || isMobile.Windows());
        }
    };

    return {
        clearFileInput: function (ctrl) {
            if (typeof jQuery != 'undefined') {
                if (ctrl instanceof jQuery) {
                    if (ctrl.length > 0) {
                        ctrl = ctrl[0];
                    }
                }
            }
            try {
                ctrl.value = null;
            } catch (ex) { }
            if (ctrl.value) {
                ctrl.parentNode.replaceChild(ctrl.cloneNode(true), ctrl);
                return false;
            }
            return true;
        },
        isMobile: isMobile,
        truncateChars: function (str, length, ending) {
            if (str) {
                if (length == null) {
                    length = 100;
                }
                if (ending == null) {
                    ending = '...';
                }
                if (str.length > length) {
                    return str.substring(0, length - ending.length) + ending;
                } else {
                    return str;
                }
            }
        },
        breakpoints: {
            xs: 0,
            sm: 544,
            md: 768,
            lg: 992,
            xl: 1200
        },
        serializedArrayToObject: function (array) {
            if (array && array instanceof Array) {
                var obj = {};
                for (var i = 0; i < array.length; i++) {
                    if (typeof (array[i]) == "object") {
                        if (array[i].hasOwnProperty('name') && array[i].hasOwnProperty('value')) {
                            obj[array[i]['name']] = array[i]['value'];
                        }
                    }
                }
                return obj;
            }
            return {};
        }
    };
});