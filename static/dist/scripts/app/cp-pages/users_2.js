function DeleteUser(e){require(["jquery","requestjs","alertjs"],function(n,t,o){var s=(new t).init(_APIURL+"/users/"+e);s.onFailure(function(e){o.alert(e.status,e.message)}),s.onSuccess(function(){document.location.reload()}),s.onComplete(function(){popupjs.hide()}),s.delete();var u=n("#alert");u.length>0&&u.css({top:"15px"})})}function UnBlockUser(e){require(["jquery","requestjs","alertjs"],function(n,t,o){var s=(new t).init(_APIURL+"/users/unblock/"+e);s.onFailure(function(e){o.alert(e.status,e.message)}),s.onSuccess(function(){document.location.reload()}),s.onComplete(function(){popupjs.hide()}),s.post();var u=n("#alert");u.length>0&&u.css({top:"15px"})})}