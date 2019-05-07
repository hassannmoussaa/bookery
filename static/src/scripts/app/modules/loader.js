define(['jquery'], function ($) {
    return {
        isInitialized: false,
        loader: null,
        init: function () {
            this.loader = $('#loader');
            this.isInitialized = true;
        },
        show: function () {
            if (!this.isInitialized) {
                this.init();
            }
            if (this.loader && this.loader.length > 0) {
                this.loader.fadeIn(300);
            }
        },
        hide: function () {
            if (!this.isInitialized) {
                this.init();
            }
            if (this.loader && this.loader.length > 0) {
                this.loader.fadeOut(300);
            }
        }
    };
});