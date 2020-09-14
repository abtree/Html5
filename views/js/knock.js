function knock(ele1, ele2){
    var b1Left = ele1.offsetLeft;
    var b1Right = b1Left + ele1.offsetWidth;
    var b1Top = ele1.offsetTop;
    var b1Bottom = b1Top + ele1.offsetHeight;
    var b2Left = ele2.offsetLeft;
    var b2Right = b2Left + ele2.offsetWidth;
    var b2Top = ele2.offsetTop;
    var b2Bottom = b2Top + ele2.offsetHeight;
    
    //console.log(`${b2Top},${b2Right},${b2Bottom},${b2Left}`);

    if (b2Right < b1Left || b1Right < b2Left || b2Bottom < b1Top || b1Bottom < b2Top) {
        return false;   //没有碰撞
    } else {
        return true;    //碰撞
    }
}