require(["jquery","royalslider"],function(a){var e=a("#embedded-album");e.length>0&&e.royalSlider({fullscreen:{enabled:!0},controlNavigation:"thumbnails",autoScaleSlider:!0,autoScaleSliderWidth:400,autoScaleSliderHeight:300,loop:!1,imageScaleMode:"fill",navigateByClick:!0,numImagesToPreload:2,arrowsNav:!0,arrowsNavAutoHide:!0,arrowsNavHideOnTouch:!0,keyboardNavEnabled:!0,fadeinLoadedSlide:!0,globalCaption:!0,globalCaptionInside:!1,thumbs:{appendSpan:!0,firstMargin:!0,paddingBottom:4,spacing:2}})});