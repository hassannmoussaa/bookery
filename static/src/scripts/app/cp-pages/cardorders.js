	function DeleteCardOrder(cardorderid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var deleteCardOrderRequest = new requestjs().init(_APIURL + '/admin/cardorders/' + cardorderid);
                deleteCardOrderRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                deleteCardOrderRequest.onSuccess(function() {
                    document.location.reload();
                });
                deleteCardOrderRequest.onComplete(function() {
                    popupjs.hide();
                });
                deleteCardOrderRequest.delete();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}


