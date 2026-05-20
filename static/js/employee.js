// OPEN MODAL
function openEmployeeModal(){

    resetEmployeeForm();

    document.getElementById("employeeModal")
    .classList.add("active");
}

// CLOSE MODAL
function closeEmployeeModal(){

    document.getElementById("employeeModal")
    .classList.remove("active");

    resetEmployeeForm();
}

function resetEmployeeForm(){

    employeeForm.reset();

    document.getElementById("employee_id").value = "";

    employeeForm.action =
    "/employees/create";

    document.querySelector(
        ".modal-header h3"
    ).innerText = "Add Employee";

    document.getElementById(
        "saveEmployeeBtn"
    ).innerText = "Save Employee";
}

function editEmployee(id){

    fetch(`/employees/json/${id}`)
    .then(res => res.json())
    .then(data => {

        document.getElementById("employee_id")
        .value = data.ID || "";

        document.getElementById("first_name")
        .value = data.FirstName || "";

        document.getElementById("middle_name")
        .value = data.MiddleName || "";

        document.getElementById("last_name")
        .value = data.LastName || "";

        document.getElementById("gender")
        .value = data.Gender || "";

        document.getElementById("religion")
        .value = data.Religion || "";

        document.getElementById("birth_place")
        .value = data.BirthPlace || "";

        if(data.BirthDate){
            document.getElementById("birth_date")
            .value = data.BirthDate;
        }

        document.getElementById("email")
        .value = data.Email || "";

        document.getElementById("phone")
        .value = data.Phone || "";

        document.getElementById("address")
        .value = data.Address || "";

        document.getElementById("employee_code")
        .value = data.EmployeeCode || "";

        document.getElementById("position")
        .value = data.Position || "";

        document.getElementById("department")
        .value = data.Department || "";

        document.getElementById("employment_status")
        .value = data.EmploymentStatus || "";

        if(data.JoinDate){
            document.getElementById("join_date")
            .value = data.JoinDate.split("T")[0];
        }

        employeeForm.action =
        `/employees/update/${id}`;

        document.querySelector(
            ".modal-header h3"
        ).innerText = "Edit Employee";

        document.getElementById(
            "saveEmployeeBtn"
        ).innerText = "Update Employee";

        document.getElementById("employeeModal")
        .classList.add("active");
    });
}

// AUTO SAVE DRAFT
const employeeForm =
document.getElementById("employeeForm");

const draftKey = "employeeDraft";

// LOAD DRAFT
window.addEventListener("DOMContentLoaded", ()=>{

    const savedDraft =
    localStorage.getItem(draftKey);

    if(savedDraft){

        const data = JSON.parse(savedDraft);

        Object.keys(data).forEach(key=>{

            const field =
            document.getElementById(key);

            if(field){
                field.value = data[key];
            }
        });
    }
});

// SAVE EVERY INPUT
employeeForm.addEventListener("input", ()=>{

    const formData = {};

    const fields =
    employeeForm.querySelectorAll(
        "input, select, textarea"
    );

    fields.forEach(field=>{

        formData[field.id] = field.value;
    });

    localStorage.setItem(
        draftKey,
        JSON.stringify(formData)
    );
});

// SUBMIT LOADING
employeeForm.addEventListener("submit", ()=>{

    const btn =
    document.getElementById("saveEmployeeBtn");

    btn.classList.add("loading");

    btn.innerText = "Saving...";

    // CLEAR DRAFT
    localStorage.removeItem(draftKey);
});

