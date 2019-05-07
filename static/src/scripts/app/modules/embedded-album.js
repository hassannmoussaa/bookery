require(['jquery', 'royalslider'], function ($) {
    var embeddedAlbum = $('#embedded-album');
    if (embeddedAlbum.length > 0) {
        embeddedAlbum.royalSlider({
            fullscreen: {
                enabled: true
            },
            controlNavigation: 'thumbnails',
            autoScaleSlider: true,
            autoScaleSliderWidth: 400,
            autoScaleSliderHeight: 300,
            loop: false,
            imageScaleMode: 'fill',
            navigateByClick: true,
            numImagesToPreload: 2,
            arrowsNav: true,
            arrowsNavAutoHide: true,
            arrowsNavHideOnTouch: true,
            keyboardNavEnabled: true,
            fadeinLoadedSlide: true,
            globalCaption: true,
            globalCaptionInside: false,
            thumbs: {
                appendSpan: true,
                firstMargin: true,
                paddingBottom: 4,
                spacing: 2
            }
        });
    }
});