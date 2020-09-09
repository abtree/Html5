function darg(obj){
    obj.onmousedown = function(evt) {
        var e = evt || window.event;
        var ox = e.offsetX;
        var oy = e.offsetY;
        window.onmousemove = function(evt){
            var e = evt || window.event;
            //obj.innerHTML = e.clientX + "," + ox + "|" + e.clientY + "," + oy;
            obj.style.left = (e.clientX - ox) + 'px';
            obj.style.top = (e.clientY - oy) + 'px';
        }
    }
    window.onmouseup = function(){
        //obj.innerHTML = "";
        window.onmousemove = null;
    }
}

function dargLimit(obj){
    obj.onmousedown = function(evt) {
        var e = evt || window.event;
        var ox = e.offsetX;
        var oy = e.offsetY;
        var wx = document.documentElement.clientWidth || document.body.clientWidth;
        var wy = document.documentElement.clientHeight || document.body.clientHeight;
        //alert(obj.clientWidth +"," + obj.clientHeight);
        window.onmousemove = function(evt){
            var e = evt || window.event;
            //obj.innerHTML = e.clientX + "," + ox + "|" + e.clientY + "," + oy;
            var l = e.clientX - ox;
            if(l < 0)
                l = 0;
            if(l > wx - obj.clientWidth)
                l = wx - obj.clientWidth;
            var t = e.clientY - oy;
            if(t < 0)
                t = 0;
            if(t > wy - obj.clientHeight)
                t = wy - obj.clientHeight;  
            obj.style.left = l + 'px';
            obj.style.top = t + 'px';
        }
    }
    window.onmouseup = function(){
        //obj.innerHTML = "";
        window.onmousemove = null;
    }
}