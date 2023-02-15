//var $ = jQuery;
var $ = jQuery.noConflict();
$(document).ready(function(){
    alert("hola");
    listado();
});

function listado(){
    $.ajax({
        url:"http://localhost:3000/platos/",
        type:"get",
        dataType:"json",
        success:function(response){
            //console.log(response);
            
            const data =  response
            console.log(data.Message);
            console.log(data.Platos);
/*             if(data.message == "Success"){
                
            }else{
                alert("predios no encontrados")
            } */
        },
        error: function(error){
            console.log(error);
        }
    });
} 
