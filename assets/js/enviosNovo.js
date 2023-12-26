// novo codigo js para conectar a api

// Event listener para o envio do formulário de busca
document.getElementById('trackingForm').addEventListener('submit', function(e) {
    e.preventDefault();
  
    // Simulação de dados de relatório de rastreamento
    const trackingCode = document.getElementById('trackingCode').value;
    const report = `
      <p><strong>Código:</strong> ${trackingCode}</p>
      <p><strong>Status:</strong> Em trânsito</p>
      <p><strong>Local:</strong> Agência de origem</p>
      <!-- Mais informações conforme necessário -->
    `;
  
    // Exibindo as informações na área designada
    document.getElementById('report').innerHTML = report;
    document.getElementById('trackingInfo').classList.remove('hidden');
  });