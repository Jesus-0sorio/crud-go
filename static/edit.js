document.addEventListener("DOMContentLoaded", async () => {
    const urlActual = window.location.pathname
    const response = await fetch('http://localhost:3000/people/'+urlActual[urlActual.length-1])
    const data = await response.json()
            document.querySelector(".container").innerHTML = `<form action="/person/edit/${data.id}" method="POST" class="m-2">
            <div class="container">
              <div class="row">
                <div class="col">
                  <div class="mb-3">
                      <label for="idContact" class="form-label"
                        >Id</label
                      >
                      ${data.id}
                      
                  </div>
                  <div class="mb-3">
                    <label for="nameContact" class="form-label">Nombre</label>
                    <input
                      type="text"
                      class="form-control"
                      name="firstname"
                      id="firstname"
                      aria-describedby="helpId"
                      placeholder=""
                      value=${data.firstname}
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
                      value=${data.lastname}
                    />
                  </div>
                </div>
                <div class="col">
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
                  <div class="mb-3">
                    <label for="cityContact" class="form-label">Ciudad</label>
                    <input
                      type="text"
                      class="form-control"
                      name="city"
                      id="city"
                      aria-describedby="helpId"
                      placeholder=""
                      value=${data.location.city}
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
                        value=${data.contact.prefix}
                      />
                  </div>
                </div>
                <div class="col">
                  <div class="mb-3">
                    <label for="numberContact" class="form-label">Numero</label>
                    <input
                      type="number"
                      class="form-control"
                      name="number"
                      id="number"
                      aria-describedby="helpId"
                      placeholder=""
                      value=${data.contact.number}
                    />
                  </div>
                  <div class="mb-3">
                    <label for="emailContact" class="form-label">Email</label>
                    <input
                      type="email"
                      class="form-control"
                      name="email"
                      id="email"
                      aria-describedby="helpId"
                      placeholder=""
                      value=${data.contact.email}
                    />
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button
                type="button"
                class="btn btn-secondary"
                data-bs-dismiss="modal"
              >
                Cerrar
              </button>
              <button type="submit" class="btn btn-primary">Guardar</button>
            </div>
          </form>`
    });
