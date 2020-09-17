function ajax({method = "get", 
               url, 
               params, 
               callback,
               errback=function(stat){
                    alert(`status=${stat}`);    
               }
            }){
    var xhttp = null;
    if (window.XMLHttpRequest) {
        // code for modern browsers
        xhttp = new XMLHttpRequest();
    } else {
        // code for old IE browsers
        xhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }

    xhttp.onreadystatechange = function() {
        if(this.readyState != 4)
            return;
        if (this.status == 200) {
            // console.log("callback:" + this.responseText);
            callback(this.responseText);
        }else{
            errback(this.status);
        }
    };
    var ps = "";
    for(var key in params){
        ps += key+"="+params[key]+"&";
    }
    if(ps!=""){
        ps = ps.substr(0, ps.length - 1);
    }
    // console.log(ps);

    if(/^get$/i.test(method)){
        // console.log("get:"+url+"?"+params);
        xhttp.open("GET", url+"?"+ps, true);
        xhttp.send();
    }else{
        // console.log("post:"+url+"?"+params);
        xhttp.open("POST", url, true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");    //这个必须写
        xhttp.send(ps);
    }
}