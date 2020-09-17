function SetCookie(name, value, {expires, path, domain, secure}){
    var c = encodeURIComponent(name)+"="+encodeURIComponent(value);
    if(expires){
        c+=";expires="+ModifyDate(expires);
    }
    if(path){
        c+=";path="+path;
    }
    if(domain){
        c+=";domain="+domain;
    }
    if(secure){
        c+=";secure"
    }
    document.cookie = c;
}

function ModifyDate(days){
    var cur = new Date();
    var day = cur.getDate();
    cur.setDate(day + days);
    return cur; 
}

function GetCookie(cname){
    var name = encodeURIComponent(cname) + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
         return decodeURIComponent(c.substring(name.length, c.length));
      }
    }
    return "";
}

function ClearCookie(name){
    document.cookie = encodeURIComponent(name)+"=;expires="+new Date(0);
}