/*
 必须遵从AMD规范 
 */
 define(function(){
     function add(x, y){
         return x + y;
     }
     function show(){
         console.log("hello world");
     }
     return{
         Add : add,
         Show : show
     }
 });