function getChar() {
    fetch("http://localhost:8080/api/characterslista").then(obj => {
    obj.json().then(json => {
        let tableContent = '';

        for (let i = 0; i < json.length; i++) {
            tableContent += `
                <tr>
                    <th>${json[i].id}</th>
                    <td>${json[i].name}</td>
                    <td><img class="image is-64x64" src="${json[i].photo}"></td>
                    <th><button class="button is-info" onclick="getCharDet(${json[i].id})">Detalles</button></th>
                </tr>
            `
        }

        document.getElementById("table-container").innerHTML = tableContent;
    })
    }).catch(err => {
        document.getElementById("not").classList.remove("hidd", "is-success")
        document.getElementById("not").classList.add("is-danger")
        document.getElementById("not-text").innerHTML=err
        console.log(err)
    })
}

function getCharDet(id) {
    fetch(`http://localhost:8080/api/characters/${id}`).then(obj => {
    obj.json().then(json => {
        document.getElementById("char-det").innerHTML = `
            <div class="container char-det-cont">
                <img class="image det-img" src="${json.imageUrl}">
                <div>
                    <p class="subtitle"><strong>Nombre :</strong> ${json.name}</p>
                    <p class="subtitle"><strong>Genero :</strong> ${json.gender}</p>
                    <p class="subtitle"><strong>Localizacion:</strong> ${json.locationName}</p>
                    <p class="subtitle"><strong>Origen :</strong> ${json.originName}</p>
                    <p class="subtitle"><strong>Especie :</strong> ${json.species}</p>
                    <p class="subtitle"><strong>Estado :</strong> ${json.status}</p>
                </div>
            </div>
        `
    })
    }).catch(err => {
        document.getElementById("not").classList.remove("hidd", "is-success")
        document.getElementById("not").classList.add("is-danger")
        document.getElementById("not-text").innerHTML=err
        console.log(err)
    })
}

function delChar() {
    fetch ("http://localhost:8080/api/characters", {
        method: "DELETE"
    }).then(obj => {
        obj.json().then(json => {
            document.getElementById("not").classList.remove("hidd", "is-danger")
            document.getElementById("not").classList.add("is-success")
            document.getElementById("not-text").innerHTML=json.message
        })
    })
    .catch(err => {
        document.getElementById("not").classList.remove("hidd", "is-success")
        document.getElementById("not").classList.add("is-danger")
        document.getElementById("not-text").innerHTML=err
        console.log(err)
    })
}

function syncChar() {
    fetch ("http://localhost:8080/api/sync/characters", {
        method: "POST"
    }).then(obj => {
        obj.json().then(json => {
            document.getElementById("not").classList.remove("hidd", "is-danger")
            document.getElementById("not").classList.add("is-success")
            document.getElementById("not-text").innerHTML=json.message
        })
    })
    .catch(err => {
        document.getElementById("not").classList.remove("hidd", "is-success")
        document.getElementById("not").classList.add("is-danger")
        document.getElementById("not-text").innerHTML=err
        console.log(err)
    })
}

function closeNot() {
    document.getElementById("not").classList.add("hidd")
}