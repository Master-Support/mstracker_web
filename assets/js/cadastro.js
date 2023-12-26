
// Event listener para enviar o formulário de cadastro
document.getElementById('objectForm').addEventListener('submit', function(e) {
  e.preventDefault();

  // Simulação de envio dos dados para a base de dados (simulado no exemplo)
  const objectName = document.getElementById('obejto').value;
  const objectCode = document.getElementById('codobjeto').value;
  // Aqui você enviaria os dados para a base de dados (simulado no exemplo)
  console.log('Objeto enviado para a base de dados:', objectName, objectCode);

  // Exibir mensagem de sucesso
  document.getElementById('successMessage').classList.remove('hidden');
});
