require(["jquery"],function(o){o(document).on("click",".dropbtn",function(n){n.preventDefault(),n.stopPropagation();var t=o(this).parent(".dropdown").find(".dropdown-content");if(t.length>0){var s=o(".dropdown-content");t.hasClass("show")?s.removeClass("show"):(s.removeClass("show"),t.addClass("show"))}}),o(window).on("click touchstart",function(n){if(!o(n.target).hasClass("dropbtn")){o(".dropdown-content").each(function(n,t){o(t).hasClass("show")&&o(t).removeClass("show")})}}),o(document).on("touchstart",".dropdown-content a",function(o){o.stopPropagation()})});