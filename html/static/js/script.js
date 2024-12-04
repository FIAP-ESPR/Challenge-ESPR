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


document.addEventListener("DOMContentLoaded", () => {
  const carList = document.getElementById("carList");
  const addCarForm = document.getElementById("addCarForm");
  let carToDelete = null;

  // Adicionar Carro
  addCarForm.addEventListener("submit", (e) => {
    e.preventDefault();

    const plate = document.getElementById("plate").value;
    const model = document.getElementById("model").value;
    const issue = document.getElementById("issue").value;

    const newCard = document.createElement("div");
    newCard.classList.add("col");
    newCard.innerHTML = `
          <div class="card h-100">
              <img src="https://via.placeholder.com/150" class="card-img-top" alt="Carro">
              <div class="card-body">
                  <h5 class="card-title">Placa: ${plate}</h5>
                  <p class="card-text">Status: Aguardando concerto.</p>
              </div>
              <div class="card-hover">
                  <h5>Modelo: ${model}</h5>
                  <p>Falha: ${issue}</p>
                  <button class="btn btn-danger btn-sm btn-delete" data-bs-toggle="modal" data-bs-target="#deleteCarModal">Excluir</button>
              </div>
          </div>
      `;

    // Adicionar evento ao botão de exclusão
    newCard.querySelector(".btn-delete").addEventListener("click", () => {
      carToDelete = newCard;
    });

    carList.querySelector(".row").appendChild(newCard);

    // Resetar formulário e fechar modal
    addCarForm.reset();
    const modal = bootstrap.Modal.getInstance(document.getElementById("addCarModal"));
    modal.hide();
  });

  // Confirmar Exclusão de Carro
  document.getElementById("confirmDelete").addEventListener("click", () => {
    if (carToDelete) {
      carToDelete.remove();
      carToDelete = null;
    }

    const modal = bootstrap.Modal.getInstance(document.getElementById("deleteCarModal"));
    modal.hide();
  });
});


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
  if (q.lenght > 3) {
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

  console.log(email, pswrd)
  if (email == testEmail && pswrd == testePswr) {
    window.location.pathname = src = 'dashboard';
    validade = true;

  }
  else {
    cadast.classList.replace('ancor-blue', 'bg-success')
  }
}

function showSignup() {
  let signup = document.querySelector('#signup');
  let login = document.querySelector('#login');

  login.classList.toggle('visually-hidden')
  signup.classList.toggle('visually-hidden')
}

