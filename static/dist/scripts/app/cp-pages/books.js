function DeleteBook(e){require(["jquery","requestjs","alertjs"],function(o,n,t){var s=(new n).init(_APIURL+"/admin/books/"+e);s.onFailure(function(e){t.alert(e.status,e.message)}),s.onSuccess(function(){document.location.reload()}),s.onComplete(function(){popupjs.hide()}),s.delete();var u=o("#alert");u.length>0&&u.css({top:"15px"})})}function VerifyBook(e){require(["jquery","requestjs","alertjs"],function(o,n,t){var s=(new n).init(_APIURL+"/book/verify/"+e);s.onFailure(function(e){t.alert(e.status,e.message)}),s.onSuccess(function(){document.location.reload()}),s.onComplete(function(){popupjs.hide()}),s.post();var u=o("#alert");u.length>0&&u.css({top:"15px"})})}function ReciveBook(e){require(["jquery","requestjs","alertjs"],function(o,n,t){var s=(new n).init(_APIURL+"/book/recive/"+e);s.onFailure(function(e){t.alert(e.status,e.message)}),s.onSuccess(function(){document.location.reload()}),s.onComplete(function(){popupjs.hide()}),s.post();var u=o("#alert");u.length>0&&u.css({top:"15px"})})}