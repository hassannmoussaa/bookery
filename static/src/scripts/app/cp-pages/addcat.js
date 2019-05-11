require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
    var addcatform = $('#add-category-form');
    if (addcatform.length > 0) {
        addcatform.on('submit', function (e) {
            e.preventDefault();
            var data = addcatform.serializeArray();
            var addCatRequest = new requestjs().init(_APIURL + '/categories', data);
            addCatRequest.onFailure(function (err) {
                alertjs.alert(err.status, err.message);
            });
            addCatRequest.onSuccess(function (data) {
                document.location = "/cp/categories";
            });
            addCatRequest.post();
        });
    }
    var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
});