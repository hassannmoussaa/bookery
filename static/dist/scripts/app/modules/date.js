require(["jquery","moment"],function(t,e){var n=t("#date");n.length>0&&setInterval(function(){n.html(e().utcOffset(3).format("ddd D MMMM YYYY hh:mm:ss a"))},1e3)});