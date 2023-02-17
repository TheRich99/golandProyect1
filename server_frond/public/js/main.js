//var $ = jQuery;
var $ = jQuery.noConflict();
$(document).ready(function(){
    
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
            if(data.Message == "Success"){

                const contenido=document.getElementById("contenido")
                
                contenido.innerHTML="";
                var contador =0;
                data.Platos.forEach(p => {
                    
                    contenido.innerHTML+=`      <section class="col">
                            <div class="card">
                            <img img id="imgMenu" src="/public/Imagenes/Menu/Churrasco.jpg" class="card-img-top">
                            <div class="card-body">
                                <h5 class="card-title">${p.nombre}</h5>
                                <p class="card-text">${p.descripcion}</p>
                                <h5 class="card-text lead"><b>$${p.precio}</b></h5>
                            </div>
                            <div class="card-footer">
                                <div class="d-grid gap-2 d-md-block">
                                <button type="button" class="btn btn-success btn-sm"><i class="fas fa-shopping-bag fa-fw"></i>
                                    Agregar</button>
                                <button type="button" class="btn btn-danger btn-sm"><i class="far fa-trash-alt"></i> Eliminar</button>
                                </div>
                            </div>
                            </div>
                        </section>`
                    
                });

                
                
            }else{
                alert("platos no encontrados")
            }
        },
        error: function(error){
            console.log(error);
        }
    });
} 

function abrir_creacion(url){
    $('#crear_plato').load(url,function(){
        $(this).modal('show');
    });
}

function cerrar_creacion(){
    $('#crear_plato').modal('hide');
}

function crear_p(){

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
        cache:false,
        processData:false,
        contentType:false,
        success: function(response){
            console.log(response);
           
            alert('Entidad registrado correctamente');
            listado();
            cerrar_creacion();

        },
        error:function(error){
            console.log(error);
            
        }

    });

}


