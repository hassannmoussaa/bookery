require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
    var loginForm = $('#login-form');
    if (loginForm.length > 0) {
        loginForm.on('submit', function (e) {
            e.preventDefault();
            var data = loginForm.serializeArray();
            var loginRequest = new requestjs().init(_APIURL + '/admins/auth', data);
            loginRequest.onFailure(function (err) {
                alertjs.alert(err.status, err.message);
            });
            loginRequest.onSuccess(function (data) {
                document.location = "/cp";
            });
            loginRequest.post();
        });
    }
    var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
});