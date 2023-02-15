$(document).ready(function() {
    $("#mostrarContra").click(function() {
      var input = $("#password");
      if (input.attr("type") === "password") {
        input.attr("type", "text");
      } else {
        input.attr("type", "password");
      }
    });
  });
  