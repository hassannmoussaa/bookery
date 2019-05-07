require(['jquery'], function ($) {
    var imgBox = $('#img-box');
    var imgBoxContent = imgBox.find('.content');
    if (imgBox.length > 0 && imgBoxContent.length > 0) {
        $('body').on('click', '.img-box', function (e) {
            e.preventDefault();
            var imgSrc = $(this).data('original-pic');
            var imgCaption = $(this).data('caption');
            var content = '<img src="';
            content += imgSrc + '" class="img" />';
            content += '<br />';
            content += '<span class="caption">' + imgCaption + '</span>';
            imgBoxContent.html(content);
            imgBox.fadeIn();
        });

        imgBox.on('click', function (e) {
            e.preventDefault();
            if (!$(e.target).hasClass('img')) {
                $(this).fadeOut();
            }
        });
    }
});