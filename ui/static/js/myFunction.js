function logout() {
    window.location.href ='/logout';
}

function addPlane(){
    window.location.href ='/flight';
}

function getUsers(){
    window.location.href ='/users';
}

function buyTicket(id){
    window.location.href ='/ticket/' + id;
}

function changeFlight(id){
    window.location.href ='/flight/' + id;
}

function sortTable(n) {
    var table, rows, switching, i, x, y, shouldSwitch, dir, switchcount = 0;
    table = document.getElementById("table");
    switching = true;
    dir = "asc";
    while (switching) {
        switching = false;
        rows = table.getElementsByTagName("tr");
        for (i = 1; i < (rows.length - 1); i++) {
            shouldSwitch = false;
            x = rows[i].getElementsByTagName("td")[n];
            y = rows[i + 1].getElementsByTagName("td")[n];
            if (dir === "asc") {
                if(n > 0 && n < 6) {
                    if (x.innerHTML.toLowerCase() > y.innerHTML.toLowerCase()) {
                        shouldSwitch = true;
                        break;
                    }
                } else if (Number(x.innerHTML) > Number(y.innerHTML)) {
                    shouldSwitch = true;
                    break;
                }

            } else if (dir === "desc") {
                if(n > 0 && n < 6) {
                    if (x.innerHTML.toLowerCase() < y.innerHTML.toLowerCase()) {
                        shouldSwitch = true;
                        break;
                    }
                } else if (Number(x.innerHTML) < Number(y.innerHTML)) {
                    shouldSwitch = true;
                    break;
                }
            }
        }
        if (shouldSwitch) {
            rows[i].parentNode.insertBefore(rows[i + 1], rows[i]);
            switching = true;
            switchcount ++;
        } else {
            if (switchcount === 0 && dir === "asc") {
                dir = "desc";
                switching = true;
            }
        }
    }
}

function cost(econom_cost, business_cost){
    document.getElementById("cost").innerHTML = 'Итоговая стоимость - ';
    document.getElementById("cost").innerHTML += parseInt(document.getElementById("econom_cost").value) * econom_cost
        + parseInt(document.getElementById("business_cost").value) * business_cost + '$';

    if(parseInt(document.getElementById("econom_cost").value) + parseInt(document.getElementById("business_cost").value) > 0){
        document.getElementById("btn").removeAttribute('disabled');
    } else {
        document.getElementById("btn").setAttribute('disabled', true);
    }
}

function maxCapacity(){
    capacityAll = document.getElementById("capacity");
    businessCapacity = document.getElementById("business_capacity");
    economCapacity = document.getElementById("econom_capacity");

    businessCapacity.setAttribute("max", parseInt(capacityAll.value) - parseInt(economCapacity.value));
    economCapacity.setAttribute("max", parseInt(capacityAll.value) - parseInt(businessCapacity.value));

    if(parseInt(businessCapacity.value) > parseInt(capacityAll.value)){
        businessCapacity.value = parseInt(capacityAll.value) - parseInt(economCapacity.value);
    }

    if(parseInt(economCapacity.value) > parseInt(capacityAll.value)){
        economCapacity.value = parseInt(capacityAll.value) - parseInt(businessCapacity.value);
    }

    if(parseInt(economCapacity.value) + parseInt(businessCapacity.value) === 0){
        capacityAll.setAttribute("min", 1);
    } else {
        capacityAll.setAttribute("min", parseInt(economCapacity.value) + parseInt(businessCapacity.value));
    }
}

function deleteFlight(id){
    var del = confirm("Вы уверены?");

    if (del === true){
        event.preventDefault();
        $.ajax({
            url: "/api/flights/" + id,
            type: "DELETE",

            success: function(data){
                window.location.href = "/";
                alert("Удаление прошло успешно");
            }
        });
    }
}

$('#update_form').submit(function(e){
    e.preventDefault();
    var formData = $(this).serialize();
    var formAction = $(this).attr('action');
    $.ajax({
        url: formAction,
        data: formData,
        type: "PUT",

        success: function(data){
            window.location.href = "/";
            alert("Изменение прошло успешно");
        }
    });
});

$('#update_ticket_form').submit(function(e){
    var ask = confirm(document.getElementById("cost").textContent + ". Вы уверены?");
    e.preventDefault();

    if (ask === true){
        var formData = $(this).serialize();
        var formAction = $(this).attr('action');
        $.ajax({
            url: formAction,
            data: formData,
            type: "PATCH",

            success: function(data){
                window.location.href = "/";
                alert("Покупка прошла успешно");
            }
        });
    }
});







