
const recibir = document.getElementById("back");

recibir.addEventListener('click', recibirDatos)

async function recibirDatos() {
    const response = await fetch('http://localhost:3000/people')
    const data = await response.json()

    console.log(data)
}