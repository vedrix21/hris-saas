// OPEN MODAL
function openEmployeeModal(){

    document.getElementById("employeeModal")
    .classList.add("active");
}

// CLOSE MODAL
function closeEmployeeModal(){

    document.getElementById("employeeModal")
    .classList.remove("active");
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