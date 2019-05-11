	function DeleteUser(userid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var deleteUserRequest = new requestjs().init(_APIURL + '/users/' + userid);
                deleteUserRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                deleteUserRequest.onSuccess(function() {
                    document.location.reload();
                });
                deleteUserRequest.onComplete(function() {
                    popupjs.hide();
                });
                deleteUserRequest.delete();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}
	function UnBlockUser(userid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var UnBlockUserRequest = new requestjs().init(_APIURL + '/users/unblock/' + userid);
                UnBlockUserRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                UnBlockUserRequest.onSuccess(function() {
                    document.location.reload();
                });
                UnBlockUserRequest.onComplete(function() {
                    popupjs.hide();
                });
                UnBlockUserRequest.post();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}

