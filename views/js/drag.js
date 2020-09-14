function Drag(node){
    //console.log(node.innerHTML);
    this.ox = 0;
    this.oy = 0;
    this.obj = node;
    var _this = this;
    this.obj.onmousedown = function (evt){
        _this.MouseDown(evt);
    };
    window.onmouseup = this.MouseUp;
    //console.log(this.obj.onmousedown);
    //console.log(window.onmouseup);
}

Drag.prototype = {
    MouseDown: function(evt){
        console.log("MouseDown");
        var e = evt || window.event;
        this.ox = e.offsetX;
        this.oy = e.offsetY;
        window.onmousemove = this.MouseMove.bind(this);
    },
    MouseMove : function(evt){
        console.log("MouseMove");
        var e = evt || window.event;
        this.obj.style.left = (e.clientX - this.ox) + 'px';
        this.obj.style.top = (e.clientY - this.oy) + 'px';
    },
    MouseUp: function(){
        console.log("MouseUp");
        window.onmousemove = null;
    }
}

function DragLimit(obj){
    Drag.call(this, obj);
    this.wx = document.documentElement.clientWidth || document.body.clientWidth;
    this.wy = document.documentElement.clientHeight || document.body.clientHeight;
}

for(var funName in Drag.prototype){
    DragLimit.prototype[funName] = Drag.prototype[funName];
    //console.log(funName);
}

DragLimit.prototype.MouseMove = function(evt){
    console.log("MouseMoveLimit");
    var e = evt || window.event;
    var l = e.clientX - this.ox;
    if(l < 0)
        l = 0;
    if(l > this.wx - this.obj.clientWidth)
        l = this.wx - this.obj.clientWidth;
    console.log(`MouseMoveLimit, l=${l}, wx=${this.wx}, clientWidth=${this.obj.clientWidth}`);
    var t = e.clientY - this.oy;
    if(t < 0)
        t = 0;
    if(t > this.wy - this.obj.clientHeight)
        t = this.wy - this.obj.clientHeight;  
    this.obj.style.left = l + 'px';
    this.obj.style.top = t + 'px';
};

function DragLimitFilter(node, func){
    DragLimit.call(this, node);
    this.func = func;
}

for(var funName in DragLimit.prototype){
    DragLimitFilter.prototype[funName] = DragLimit.prototype[funName];
    //console.log(funName);
}

DragLimitFilter.prototype.MouseMove = function(evt){
    DragLimit.prototype.MouseMove.call(this, evt);
    this.func();
}

// function darg(obj){
//     obj.onmousedown = function(evt) {
//         var e = evt || window.event;
//         var ox = e.offsetX;
//         var oy = e.offsetY;
//         window.onmousemove = function(evt){
//             var e = evt || window.event;
//             //obj.innerHTML = e.clientX + "," + ox + "|" + e.clientY + "," + oy;
//             obj.style.left = (e.clientX - ox) + 'px';
//             obj.style.top = (e.clientY - oy) + 'px';
//         }
//     }
//     window.onmouseup = function(){
//         //obj.innerHTML = "";
//         window.onmousemove = null;
//     }
// }

// function dargLimit(obj){
//     obj.onmousedown = function(evt) {
//         var e = evt || window.event;
//         var ox = e.offsetX;
//         var oy = e.offsetY;
//         var wx = document.documentElement.clientWidth || document.body.clientWidth;
//         var wy = document.documentElement.clientHeight || document.body.clientHeight;
//         //alert(obj.clientWidth +"," + obj.clientHeight);
//         window.onmousemove = function(evt){
//             var e = evt || window.event;
//             //obj.innerHTML = e.clientX + "," + ox + "|" + e.clientY + "," + oy;
//             var l = e.clientX - ox;
//             if(l < 0)
//                 l = 0;
//             if(l > wx - obj.clientWidth)
//                 l = wx - obj.clientWidth;
//             var t = e.clientY - oy;
//             if(t < 0)
//                 t = 0;
//             if(t > wy - obj.clientHeight)
//                 t = wy - obj.clientHeight;  
//             obj.style.left = l + 'px';
//             obj.style.top = t + 'px';
//         }
//     }
//     window.onmouseup = function(){
//         //obj.innerHTML = "";
//         window.onmousemove = null;
//     }
// }

// function dargLimitFunc(obj, func){
//     obj.onmousedown = function(evt) {
//         var e = evt || window.event;
//         var ox = e.offsetX;
//         var oy = e.offsetY;
//         var wx = document.documentElement.clientWidth || document.body.clientWidth;
//         var wy = document.documentElement.clientHeight || document.body.clientHeight;
//         //alert(obj.clientWidth +"," + obj.clientHeight);
//         window.onmousemove = function(evt){
//             var e = evt || window.event;
//             //obj.innerHTML = e.clientX + "," + ox + "|" + e.clientY + "," + oy;
//             var l = e.clientX - ox;
//             if(l < 0)
//                 l = 0;
//             if(l > wx - obj.clientWidth)
//                 l = wx - obj.clientWidth;
//             var t = e.clientY - oy;
//             if(t < 0)
//                 t = 0;
//             if(t > wy - obj.clientHeight)
//                 t = wy - obj.clientHeight;  
//             obj.style.left = l + 'px';
//             obj.style.top = t + 'px';

//             func();
//         }
//     }
//     window.onmouseup = function(){
//         //obj.innerHTML = "";
//         window.onmousemove = null;
//     }
// }