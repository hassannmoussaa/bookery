require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
    var addadminform = $('#add-admin-form');
    if (addadminform.length > 0) {
        addadminform.on('submit', function (e) {
            e.preventDefault();
            var data = addadminform.serializeArray();
            var addadminRequest = new requestjs().init(_APIURL + '/admins', data);
            addadminRequest.onFailure(function (err) {
                alertjs.alert(err.status, err.message);
            });
            addadminRequest.onSuccess(function (data) {
                document.location = "/cp/admins";
            });
            addadminRequest.post();
        });
    }
    var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
});