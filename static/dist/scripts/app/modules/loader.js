define(["jquery"],function(i){return{isInitialized:!1,loader:null,init:function(){this.loader=i("#loader"),this.isInitialized=!0},show:function(){this.isInitialized||this.init(),this.loader&&this.loader.length>0&&this.loader.fadeIn(300)},hide:function(){this.isInitialized||this.init(),this.loader&&this.loader.length>0&&this.loader.fadeOut(300)}}});