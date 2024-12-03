// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('click', event => {
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }

      form.classList.add('was-validated')
    }, false)
  })
})()


const btnSend = document.getElementById('cadastrar');
btnSend.addEventListener('click', defineUser);

function defineUser() {
  //função validadar usuário editar conforme as permissoes
};



function displayCards() {
  //função para mostar Cards de acordo com as nececidades do usuário)
  let dashboard = document.querySelector('#dashboard');

};

function countOverview() {
  //funçao pega o total de: Mecanicos, Peças disponiveis, Carros em Manutençao para o Overview do Dashboard
  //let numMecanicos = pegar valores na base de dados ;
  //let numPecas = pegar valores na base de dados //;
  //let numCarros = pegar valores na base de dados ;
  let Mecanicos = document.querySelector('#res-mecanicos-n');
  let Pecas = document.querySelector('#res-pecas-n');
  let Carros = document.querySelector('#res-carros-n');
}

const quest = document.getElementById('iaInput');
quest.on('click', loadAi());

function loadAi() {
  window.alert('GEPETTo')
  // função que verifica se o usuário digitou na pesquisa e mostra animação de carregamento
  const quest = document.getElementById('iaInput')
  let spn = document.getElementById('spinner');
  var q = quest.value;
  if(q.lenght > 3 ){
    spn.classList.remove('visually-hidden')
  }
  
}


function validadeLogin() {
  //Valida o Usuário procurando Email e senha em usuários caastrados
  const validade = false;
  
  let cadast = document.getElementById('cadastrarbtn');
  let email = document.getElementById('emailv').value;
  let pswrd = document.getElementById('examplePassword').value;
  let testEmail = "cliente@email.com";
  let testePswr = "Teste@1234"

  console.log(email,pswrd)
  if (email == testEmail && pswrd == testePswr) {
    window.location.pathname = src ='dashboard';
    validade = true;

  }
  else {
    cadast.classList.replace('ancor-blue','bg-success')
  }
}

function showSignup() {
  let signup = document.querySelector('#signup');
  let login = document.querySelector('#login');

  login.classList.toggle('visually-hidden')
  signup.classList.toggle('visually-hidden')
}

