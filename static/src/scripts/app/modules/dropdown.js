require(['jquery'], function ($) {

    $(document).on('click', '.dropbtn', function (e) {
        e.preventDefault();
        e.stopPropagation();
        var dropdown = $(this).parent('.dropdown').find('.dropdown-content');
        if (dropdown.length > 0) {
            var dropdowns = $('.dropdown-content');
            if (dropdown.hasClass('show')) {
                dropdowns.removeClass('show');
            } else {
                dropdowns.removeClass('show');
                dropdown.addClass('show');
            }
        }
    });

    // Close the dropdown menu if the user clicks outside of it
    $(window).on('click touchstart', function (event) {
        if (!$(event.target).hasClass('dropbtn')) {
            var dropdowns = $('.dropdown-content');
            dropdowns.each(function (i, dropdown) {
                if ($(dropdown).hasClass('show')) {
                    $(dropdown).removeClass('show');
                }
            });
        }
    });

    $(document).on('touchstart', '.dropdown-content a', function (e) {
        e.stopPropagation();
    });
});