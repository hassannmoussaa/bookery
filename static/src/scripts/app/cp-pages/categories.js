	function DeleteCategory(catid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var deleteCategoryRequest = new requestjs().init(_APIURL + '/admin/categories/' + catid);
                deleteCategoryRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                deleteCategoryRequest.onSuccess(function() {
                    document.location.reload();
                });
                deleteCategoryRequest.onComplete(function() {
                    popupjs.hide();
                });
                deleteCategoryRequest.delete();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}

