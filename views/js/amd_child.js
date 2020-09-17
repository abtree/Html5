/*
 必须遵从AMD规范 
 */
define(["father"], function(father){
    function double(x, y, mul){
        return father.Add(x, y) * mul;
    }
    return{
        Double : double,
    }
});