$('#formulario-cadastro-objeto').on('submit', rastrearObjeto);

function rastrearObjeto(e) {
    e.preventDefault();

    $.ajax({
        url: "/adicionar-encomendas",
        method: "POST",
        data: {
            "nomeObjeto": $('#nomeObjeto').val(),
            "codigoObjeto": $('#codigoObjeto').val()
        }
    });
}

 
