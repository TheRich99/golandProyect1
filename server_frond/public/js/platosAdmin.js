function crear_p(){


      
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



    function editar_p(){


      
        const form = document.getElementById('form_editar_p');
        const formData = new FormData(form);
        console.log(form);
        $.ajax({
            
            url: $('#form_editar_p').attr('action'),
            type: $('#form_editar_p').attr('method'),
            data: formData,
            cache:false,
            processData:false,
            contentType:false,
            success: function(response){
                console.log(response);
               
                alert('Entidad registrado correctamente');
                listado();
                cerrar_edicion();
    
            },
            error:function(error){
                console.log(error);
                
            }
    
        });
    
    }

    function abrir_eliminacionP(id){
        
        var datos = { 'id' :id };
        $.ajax({
            
            url: 'http://localhost:3000/platos/',
            type: 'delete',
            data: datos,
            cache:false,
            processData:false,
            contentType:false,
            success: function(response){
                console.log(response);
               
                alert('plato eliminado correctamente');
                listado();
                /* cerrar_e(); */
    
            },
            error:function(error){
                console.log(error);
                
            }
    
        });
    }
    




    function abrir_creacion(url){
        $('#crear_plato').load(url,function(){
            $(this).modal('show');
        });
    }
    
    function cerrar_edicion(){
        $('#editar_plato').modal('hide');
    }


    var $ = jQuery.noConflict();
$(document).ready(function(){
    
    listado();
});


function abrir_edicion(url,ID,nombre,precio,descripcion){
    $('#editar_plato').load(url,function(){
        $("#ID2").val(ID);
        $("#campoNombre2").val(nombre);
        $("#precio2").val(precio);
        $("#cantidad2").val(0);
        $("#descrip2").val(descripcion);
        
        $(this).modal('show');
        
    });
}

    
function cerrar_creacion(){
    $('#crear_plato').modal('hide');
}
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

                //const contenido=document.getElementById("contenido")
                const tabla=document.getElementById('tabla_platos')
                tabla.innerHTML="";
                var contador =0;
                tabla.innerHTML+=`<tr>
                                    <th>id</th>
                                    <th>Nombre</th> 
                                    <th>Precio</th>
                                    <th>Descripci??n</th>
                                    <th>Opciones</th>
                                </tr>`;
                data.Platos.forEach(p => {
                    
                    tabla.innerHTML += `               
                    <tr>
                    <td>${p.ID}</td>
                    <td>${p.nombre}</td> 
                    <td>$${p.precio}</td>
                    <td>${p.descripcion}</td>
                    
                    <td>
                        <button onclick="abrir_edicion('/public/Platos/modal_plato_editar.html',${p.ID},'${p.nombre}',${p.precio},'${p.descripcion}')"  class="btn btn-info btn-sm">Editar</button>
                        <button onclick="abrir_eliminacionP('${p.ID}');"  class="btn btn-danger btn-sm">Eliminar</button>
                    </td>
            
            
                </tr>`
                    
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

