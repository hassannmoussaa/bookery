	function DeleteAdmin(adminid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var deleteAdminRequest = new requestjs().init(_APIURL + '/admin/' + adminid);
                deleteAdminRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                deleteAdminRequest.onSuccess(function() {
                    document.location.reload();
                });
                deleteAdminRequest.onComplete(function() {
                    popupjs.hide();
                });
                deleteAdminRequest.delete();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}

