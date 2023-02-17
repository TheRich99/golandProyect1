var noMesa = document.getElementById("noMesa");
var detalle = document.getElementById("detalle");
var total = document.getElementById("total");

var fechaP = document.getElementById("fechaP");
var platosP = document.getElementById("platosP");
var valorP = document.getElementById("valorP");
var mesaP = document.getElementById("mesaP");


function detalles(a,b,c){
    fechaP.value =  new Date(); 
    platosP.value = b;
    valorP.value = c;
    mesaP.value = a;
}
