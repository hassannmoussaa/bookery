define([], function () {
    return {
        queryParams: {
            search: null,
            action: null,
            entryId: null,
            page: null
        },
        isInitialized: false,
        initialUrl: null,
        currentUrl: null,
        onChangeHandler: null,
        init: function () {
            var url = new Url().toString();
            this.initialUrl = url;
            this.currentUrl = url;
            this.isInitialized = true;
            var self = this;
            if (window.history && window.history.pushState) {
                window.onpopstate = function (e) {
                    if (e.state) {
                        self.parseURL();
                        self.refresh(true);
                    }
                };
            }
            this.parseURL();
        },
        getQueryParam: function (key) {
            if (key && typeof key == 'string') {
                if (this.queryParams.hasOwnProperty(key)) {
                    this.parseURL();
                    return this.queryParams[key];
                }
            }
            return null;
        },
        changeSearch: function (v) {
            if (typeof v == 'string') {
                v = v.trim();
                if (v) {
                    this.queryParams.search = v;
                } else {
                    this.queryParams.search = null;
                }
            } else {
                this.queryParams.search = null;
            }
            this.queryParams.page = 1
        },
        changeEntryId: function (v) {
            if (typeof v == 'string') {
                v = v.trim();
            }
            v = Number(v);
            if (!isNaN(v) && v > 0) {
                this.queryParams.entryId = v;
            } else {
                this.queryParams.entryId = null;
            }
        },
        changePage: function (v) {
            if (typeof v == 'string') {
                v = v.trim();
            }
            v = Number(v);
            if (!isNaN(v) && v > 0) {
                this.queryParams.page = v;
            } else {
                this.queryParams.page = null;
            }
        },
        refresh: function (isBackAction, isFirstRefresh) {
            if (!this.isInitialized) {
                this.init();
            }
            var url = this.generateURL();
            if (this.currentUrl != url || !(window.history && window.history.pushState) || isFirstRefresh) {
                this.currentUrl = url;
                if (!isBackAction) {
                    if (window.history && window.history.pushState) {
                        window.history.pushState({}, "", url);
                    }
                }
                if (typeof this.onChangeHandler == 'function') {
                    this.onChangeHandler(this.queryParams, url)
                }
            }
        },
        changeAction: function (v) {
            if (typeof v == 'string') {
                v = v.trim();
                if (v) {
                    this.queryParams.action = v;
                } else {
                    this.queryParams.action = null;
                }
            } else {
                this.queryParams.action = null;
            }
        },
        generateURL: function () {
            var u = new Url();

            delete u.query.search;
            delete u.query.action;
            delete u.query.entry_id;
            delete u.query.page;

            if (typeof this.queryParams.search == 'string' && this.queryParams.search.trim() != "") {
                u.query.search = this.queryParams.search.trim();
            }
            if (typeof this.queryParams.action == 'string' && this.queryParams.action.trim() != "") {
                u.query.action = this.queryParams.action.trim().toLowerCase();
            }
            if (!isNaN(Number(this.queryParams.entryId))) {
                u.query.entry_id = this.queryParams.entryId;
            }
            if (!isNaN(Number(this.queryParams.page))) {
                u.query.page = this.queryParams.page;
            }
            return u.toString();
        },
        parseURL: function (u) {
            if (typeof u == 'undefined') {
                u = new Url();
            } else if (typeof u == 'string') {
                u = new Url(u);
            }
            this.queryParams = {
                search: null,
                action: null,
                entryId: null,
                page: null
            };
            if (typeof u.query.search == 'string' && u.query.search.trim() != "") {
                this.queryParams.search = u.query.search.trim();
            }

            if (typeof u.query.action == 'string' && u.query.action.trim() != "") {
                this.queryParams.action = u.query.action.trim().toLowerCase();
            }
            if (typeof u.query.entry_id == 'string' && u.query.entry_id.trim() != "") {
                if (!isNaN(Number(u.query.entry_id.trim()))) {
                    this.queryParams.entryId = Number(u.query.entry_id.trim());
                }
            }
            if (typeof u.query.page == 'string' && u.query.page.trim() != "") {
                if (!isNaN(Number(u.query.page.trim()))) {
                    this.queryParams.page = Number(u.query.page.trim());
                }
            }
        }
    };
});