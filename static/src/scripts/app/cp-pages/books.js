	function DeleteBook(bookid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var deleteBookRequest = new requestjs().init(_APIURL + '/admin/books/' + bookid);
                deleteBookRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                deleteBookRequest.onSuccess(function() {
                    document.location.reload();
                });
                deleteBookRequest.onComplete(function() {
                    popupjs.hide();
                });
                deleteBookRequest.delete();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}
	function VerifyBook(bookid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var VerifyBookRequest = new requestjs().init(_APIURL + '/book/verify/' + bookid);
                VerifyBookRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                VerifyBookRequest.onSuccess(function() {
                    document.location.reload();
                });
                VerifyBookRequest.onComplete(function() {
                    popupjs.hide();
                });
                VerifyBookRequest.post();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}
	function ReciveBook(bookid) { 
					 require(['jquery', 'requestjs', 'alertjs'], function ($, requestjs, alertjs) {
					  var ReciveBookRequest = new requestjs().init(_APIURL + '/book/recive/' + bookid);
                ReciveBookRequest.onFailure(function(err) {
                    alertjs.alert(err.status, err.message);
                });
                ReciveBookRequest.onSuccess(function() {
                    document.location.reload();
                });
                ReciveBookRequest.onComplete(function() {
                    popupjs.hide();
                });
                ReciveBookRequest.post();
				
				  var alertjsContainer = $('#alert');
    if (alertjsContainer.length > 0) {
        alertjsContainer.css({ top: '15px' });
    }
			   });
}

