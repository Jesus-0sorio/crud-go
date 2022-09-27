document.addEventListener("DOMContentLoaded", async () => {
    const response = await fetch('http://localhost:3000/people')
    const data = await response.json()
    data.forEach(element => {
        document.querySelector("#cardContact").innerHTML += `<div class="col">
        <div class="card text-center" style="width: 18rem;">
        <div class="card-body">
        <h5 class="card-title">${element.firstname} ${element.lastname}</h5>
        <p class="card-text">${element.location.city}, ${element.location.country} <br>   ${element.contact.prefix} ${element.contact.number} <br> ${element.contact.email}</p>
        <div class="mb-3">
        <a href="/people/delete/${element.id}"  class="btn btn-danger">Eliminar</a> | <a href="/people/edit/${element.id}"  class="btn btn-secondary">Editar</a>
        </div>
        <div class="mb-3">
        ${element.id}
        </div>
        </div>
        </div>
        </div>`
      });

    })
const recibir = document.getElementById("back")


async function recibirDatos() {
    const response = await fetch('http://localhost:3000/people')
    const data = await response.json()
    console.log(data)
}


