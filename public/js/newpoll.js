const pollOptions = document.getElementById("poll-options")

const createNewOption = () =>{
    const firstOption = pollOptions.childNodes[0];
    const newOption = document.createElement('input');
    newOption.value = "";
    newOption.placeholder = "type option here";
    newOption.name = firstOption.name;

    pollOptions.appendChild(newOption);
}


const newOptionButton = document.getElementById("new-opt");
newOptionButton.addEventListener("click", () => {
    createNewOption();
});

function submitForm() {
    
}