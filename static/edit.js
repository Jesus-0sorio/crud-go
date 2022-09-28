document.addEventListener("DOMContentLoaded", async () => {
    const urlActual = window.location.pathname
    const response = await fetch('http://localhost:3000/people/'+urlActual[urlActual.length-1])
    const data = await response.json()
    document.querySelector("#editContact").innerHTML = `<form action="/person/edit/${data.id}" method="POST" class="m-1">
    <div class="card">
      <div class="card-header">
      <h3>Edit contact</h3>
      </div>
      <div class="row p-3">
        <div class="col">
          <div class="mb-3">
            <label for="nameContact" class="form-label">Nombre</label>
            <input
              type="text"
              class="form-control"
              name="firstname"
              id="firstname"
              aria-describedby="helpId"
              placeholder=""
              value="${data.firstname}"
            />
          </div>
          <div class="mb-3">
            <label for="lastNameContact" class="form-label">Apellido</label>
            <input
              type="text"
              class="form-control"
              name="lastname"
              id="lastname"
              aria-describedby="helpId"
              placeholder=""
              value="${data.lastname}"
            />
          </div>
          <div class="mb-3">
            <label for="countryContact" class="form-label">Pais</label>
            <input
              type="text"
              class="form-control"
              name="country"
              id="country"
              aria-describedby="helpId"
              placeholder=""
              value="${data.location.country}"
            />
          </div>
        </div>
        <div class="col">
          <div class="mb-3">
            <label for="cityContact" class="form-label">Ciudad</label>
            <input
              type="text"
              class="form-control"
              name="city"
              id="city"
              aria-describedby="helpId"
              placeholder=""
              value="${data.location.city}"
            />
          </div>
          <div class="mb-3">
              <label for="prefijoContact" class="form-label">Prefijo</label>
              <input
                type="text"
                class="form-control"
                name="prefix"
                id="prefix"
                aria-describedby="helpId"
                placeholder=""
                value="${data.contact.prefix}"
              />
          </div>
          <div class="mb-3">
            <label for="numberContact" class="form-label">Numero</label>
            <input
              type="text"
              class="form-control"
              name="number"
              id="number"
              aria-describedby="helpId"
              placeholder=""
              value="${data.contact.number}"
            />
          </div>
        </div>
        <div class="col">
          <div class="mb-3">
            <label for="emailContact" class="form-label">Email</label>
            <input
              type="email"
              class="form-control"
              name="email"
              id="email"
              aria-describedby="helpId"
              placeholder=""
              value="${data.contact.email}"
            />
          </div>
        </div>
      </div>
      <div class="card-footer d-flex justify-content-end">
        <a
          class="btn btn-secondary"
          href="/"
        >
          Cerrar
        </a>
        <button type="submit" class="btn btn-primary mx-2" href="/">Guardar</button>
      </div>
    </div>
  </form>`
    });
