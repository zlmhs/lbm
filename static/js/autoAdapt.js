function resetZoom(){
    var winW = document.documentElement.clientWidth;
    var winH = document.documentElement.clientHeight;
    var wrap = document.getElementsByClassName('wrap')[0];
    var scale = winW/750;
    // var realH = winH/scale;
    document.body.style.zoom = scale;
    // wrap.style.height = realH-108 +'px';
}
window.onresize = window.onload = function () {
    resetZoom();
}
