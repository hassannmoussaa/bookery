require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
    var signupForm = $('#signup');
	
    if (signupForm.length > 0) {
        signupForm.on('submit', function (e) {
            e.preventDefault();
            var data = signupForm.serializeArray();
            var signupRequest = new requestjs().init(_APIURL + '/users', data);
            signupRequest.onFailure(function (err) {
                alertjs.alert(err.status, err.message);
            });
            signupRequest.onSuccess(function (data) {
				document.cookie = "FinishSignUp=True"
			 document.location = "/";
			
           
            });
            signupRequest.post();
        });
    }
    var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
});