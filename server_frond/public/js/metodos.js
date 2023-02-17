$(document).ready(function () {
  $("#mostrarContra").click(function () {
    var input = $("#password");
    if (input.attr("type") === "password") {
      input.attr("type", "text");
    } else {
      input.attr("type", "password");
    }
  });
});

var queryString = window.location.search;
var urlParams = new URLSearchParams(queryString);
var arregloString = urlParams.get("arregloPedido");
let arreglo = JSON.parse(decodeURIComponent(arregloString));

var listaPedido = document.getElementById("listaProductos");
var cantidadPlatos = document.getElementById("cantidadP");
var totalP = document.getElementById("total");

var total = parseInt(0);
var cantidadP = 0;

cargarPlatosPedido(arreglo);

function cargarPlatosPedido(arregloPedido) {
  for (i = 0; i < arregloPedido.length; i++) {
    listaPedido.innerHTML += `<div class="container-fluid" id="${arregloPedido[i][2]}">
                                                      <hr><br>
                                                      <h5 class="poppins-regular font-weight-bold full-box text-center" >${arregloPedido[i][0]}</h5><br>
                                                      <div class="bag-item full-box">
                                                        <div class="full-box">
                                                          <div class="container-fluid">
                                                            <div class="row">
                                                              <div class="col-12 col-lg-6 text-center mb-4">
                                                                <div class="row justify-content-center">
                                                                  <div class="col-auto">
                                                                    <div class="form-outline mb-4">
                                                                      <input type="number" value="${arregloPedido[i][1]}" class="form-control text-center" id="cantidad"
                                                                        maxlength="10" disabled>
                                                                      <label class="form-label">Cantidad</label>
                                                                    </div>
                                                                  </div>
                                                                </div>
                                                              </div>
                                                              <div class="col-12 col-lg-4 text-center mb-4">
                                                                <span class="poppins-regular font-weight-bold">SUBTOTAL: $${arregloPedido[i][3]}</span>
                                                              </div>

                                                            </div>
                                                          </div>
                                                        </div>
                                                      </div>
                                                      <hr>
                                                    </div>`;
    total += parseInt(arregloPedido[i][3]) * parseInt(arregloPedido[i][1]);
    cantidadP += arregloPedido[i][1];
  }
  cantidadPlatos.innerHTML = cantidadP;
  totalP.innerHTML = "$" + total;
}


function confirmarOrden() {
  const fecha = new Date();
  const select = document.getElementById("selectMesa");
  const valorSeleccionado = select.value;
  
  console.log(fecha, total, valorSeleccionado);
}