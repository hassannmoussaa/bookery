require(['jquery', 'requestjs', 'alertjs', 'text', 'helpers'], function ($, requestjs, alertjs, text, helpers) {
    //send message
    var uploadForm = $('#upload-form');
    if (uploadForm.length > 0) {
        uploadForm.on('submit', function (e) {
            e.preventDefault();
            var data = new FormData(this);
            var uploadRequest = new requestjs().init(_APIURL + '/upload', data, true);
            var alertId = alertjs.alert('info', text.uploading, true);
            uploadRequest.onProgress(function (percentage) {
                alertjs.updateRealProgress(alertId, percentage);
            });
            uploadRequest.onFailure(function (err) {
                alertjs.alert(err.status, err.message);
            });
            uploadRequest.onSuccess(function (result) {
                helpers.clearFileInput(uploadForm.find('#files'));
                alertjs.alert(result.status, result.message);
            });
            uploadRequest.post();
        });
    }
});