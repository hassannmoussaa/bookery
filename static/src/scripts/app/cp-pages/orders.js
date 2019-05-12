	function DeleteOrder(orderid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var deleteOrderRequest = new requestjs().init(_APIURL + '/admin/orders/' + orderid);
                deleteOrderRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                deleteOrderRequest.onSuccess(function() {
                    document.location.reload();
                });
                deleteOrderRequest.onComplete(function() {
                    popupjs.hide();
                });
                deleteOrderRequest.delete();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}

	function CompletOrder(orderid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var CompletOrderRequest = new requestjs().init(_APIURL + '/complet/orders/' + orderid);
                CompletOrderRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                CompletOrderRequest.onSuccess(function() {
                    document.location.reload();
                });
                CompletOrderRequest.onComplete(function() {
                    popupjs.hide();
                });
                CompletOrderRequest.post();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}

