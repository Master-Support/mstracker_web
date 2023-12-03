// script.js
$(document).ready(function() {
    carregarEnvios();
});


// Função para carregar os envios usando AJAX
function carregarEnvios() {
    $.ajax({
        url: '/vizualizar-encomendas', // Substitua pela rota correta em sua aplicação
        type: 'GET',
        dataType: 'json',
        success: function(response) {
            exibirEnvios(response.envios);
        },
        error: function(error) {
            console.error('Erro ao obter os envios:', error);
        }
    });
}

// Função para exibir os envios na tabela
function exibirEnvios(envios) {
    var tbody = $('#envios-table tbody');
    tbody.empty();er

    $.each(envios, function(index, envio) {
        var row = '<tr>';
        row += '<td>' + envio.nomeObjeto + '</td>';
        row += '<td>' + envio.codigoObjeto + '</td>';
        row += '<td>' + envio.dataPrevistaDeEntrega + '</td>';
        row += '<td>' + envio.statusObjeto + '</td>';
        row += '<td>' + envio.localizacao + '</td>';
        row += '</tr>';
        tbody.append(row);
    });
}
