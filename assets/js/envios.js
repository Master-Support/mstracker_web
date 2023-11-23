$('#formulario-cadastro-objeto').on('submit', rastrearObjeto);

function rastrearObjeto(e) {
    e.preventDefault();

    $.ajax({
        url: "/adicionar-encomendas",
        method: "POST",
        data: {
            "codigoObjeto": $('#nomeObjeto').val(),
            "nomeObjeto": $('#codigoObjeto').val()
        }
    });
}