function randColor(){
    parseInt(Math.random()*256);
    return "rgba(" + parseInt(Math.random()*256)+","+parseInt(Math.random()*256)+","+parseInt(Math.random()*256)+",1)";
}

function addEvent(node, evtTyp, funcName){
    if(node.addEventListener){
        node.addEventListener(evtTyp, funcName, false);
    }else{
        node.attachEvent("on" + evtTyp, funcName);
    }
}

function removeEvent(node, evtTyp, funcName){
    if(node.removeEventListener){
        node.removeEventListener(evtTyp, funcName);
    }else{
        node.detachEvent("on" + evtTyp, funcName);
    }
}