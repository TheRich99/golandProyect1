//var $ = jQuery;
var $ = jQuery.noConflict();
var inventarioPedido;

var contador = 0;
let arregloPedido = [];
var masBoton = document.getElementById("masbutton");
var menosBoton = document.getElementById("menosbutton");




function mas(boton) {
    //precio sin $
    let precioC;
    //div contenedor 
    var div = boton.closest(".card");
    // Nombre plato
    var h5 = div.querySelector(".card-title");
    var nombre = h5.textContent;
    // contador plato
    var div2 = div.querySelector(".contador");
    var contador2 = div2.textContent;
    // id plato
    var div3 = div.querySelector("#idP");
    var id = div3.textContent;
    // precio plato
    var div4 = div.querySelector("#precioP");
    var precio = div4.textContent.substring(1);

    var existe = false;
    var position;
    contador = contador2;
    contador++;

    // Si el elemento existe, actualiza su cantidad
    if (contador > 0) {
        for (let i = 0; i < arregloPedido.length; i++) {
            if (arregloPedido[i][0] === nombre) {
                position = i;
                existe = true;
                //console.log("existe," + nombre);
                break;
            }
        }
        if (existe) {
            arregloPedido[position][1] = contador;
            existe = false;
            //console.log("contador+:" + contador + "," + arregloPedido[position][0]);
        } else {
            //console.log("nuevo");
            arregloPedido.push([nombre, 1, id, precio]);
            //console.log("arreglo+" + arregloPedido)
        }
    }
    div2.innerHTML = contador;
}

function menos(boton) {
    var div = boton.closest(".card");
    var h5 = div.querySelector(".card-title");
    var nombre = h5.textContent;
    var div2 = div.querySelector(".contador");
    var contador2 = div2.textContent;

    contador = contador2;

    if (contador > 0) {
        contador--;
        //console.log("contador-:" + contador + "," + nombre);
    }
    if (contador == 0) {
        for (let i = 0; i < arregloPedido.length; i++) {
            if (arregloPedido[i][0] === nombre) {
                arregloPedido.splice(i, 1);
                break;
            }
        }
    }
    div2.innerHTML = contador;
}

function crearPedido() {
    if (arregloPedido.length === 0) {
        window.alert("asegurate de seleccionar al menos un plato");
    } else {
        var queryString = "?arregloPedido=" + encodeURIComponent(JSON.stringify(arregloPedido));
        var url = "Pedido.html" + queryString;
        window.location.href = url;
    }
}

$(document).ready(function () {

    listado();
});
function contador() {

}

function listado() {
    $.ajax({
        url: "http://localhost:3000/platos/",
        type: "get",
        dataType: "json",
        success: function (response) {
            //console.log(response);

            const data = response
            console.log(data.Message);
            console.log(data.Platos);
            if (data.Message == "Success") {

                const contenido = document.getElementById("contenido")

                contenido.innerHTML = "";
                var contador = 0;
                data.Platos.forEach(p => {

                    contenido.innerHTML += `      <section class="col">
                            <div class="card">
                                <img img id="imgMenu" src="/public/Imagenes/Menu/Churrasco.jpg" class="card-img-top">
                                <div class="card-body">
                                    <p class="card-text" id="idP" style="display:none">${p.ID}</p>
                                    <h5 class="card-title" id="nombre">${p.nombre}</h5>
                                    <p class="card-text">${p.descripcion}</p>
                                    <h5 class="card-text lead" id="precioP"><b>$${p.precio}</b></h5>
                                </div>
                                <div class="card-footer ">
                                    <div class="d-grid gap-2 d-md-block align-items-end">
                                    <button id="masbutton" onclick="mas(this)" type="button" class="btn btn-success btn-sm"><i class="fa-solid fa-plus">+</i></button>
                                    <button id="menosbutton" onclick="menos(this)" type="button" class=" btn btn-danger btn-sm"> <i class="fa-solid fa-minus">-</i></button>
                                    <div class="contador" id="contador">0</div>
                                    </div>
                                </div>
                            </div>
                        </section>`

                });



            } else {
                alert("platos no encontrados")
            }
        },
        error: function (error) {
            console.log(error);
        }
    });
}

function abrir_creacion(url) {
    $('#crear_plato').load(url, function () {
        $(this).modal('show');
    });
}

function cerrar_creacion() {
    $('#crear_plato').modal('hide');
}

function crear_p() {

    /*     alert("hola");
        let nombre = document.getElementById("campoNombre")
        let descripcion = document.getElementById("exampleFormControlTextarea1")
        let precio = document.getElementById("precio")
        let cantidad = document.getElementById("cantidad")
        let imagen = document.getElementById("imagen")
    
        var formData = new FormData();
        console.log(nombre.value);
        formData.append('nombre',nombre.value);
        formData.append('precio',precio.value);
        formData.append('cantidad',cantidad.value);
        formData.append('descripcion',descripcion.value);
        formData.append('imagen',imagen.value); */


    const form = document.getElementById('form_crear_p');
    const formData = new FormData(form);
    console.log(form);
    $.ajax({

        url: $('#form_crear_p').attr('action'),
        type: $('#form_crear_p').attr('method'),
        data: formData,
        cache: false,
        processData: false,
        contentType: false,
        success: function (response) {
            console.log(response);

            alert('Entidad registrado correctamente');
            listado();
            cerrar_creacion();

        },
        error: function (error) {
            console.log(error);

        }

    });

}


