import TableURL from "./url/table.js";
import SigninElement from "./component/signin-toolbar.js";
import FetchWrapper from "./common/fetch-wrapper.js";
import getConfigs from "./defaults.js";

customElements.define("signin-btn", SigninElement);

const table = new TableURL("#container-links");
const configs = getConfigs();
const API = new FetchWrapper(configs.apiEndpoint);

async function saveLink() {
  const request = {};
  document
    .querySelectorAll("#form-new-link input, #form-new-link textarea")
    .forEach((input) => {
      request[input.name] = input.value;
    });
  const [status, data] = await API.post("/urls", request);
  if (status === 200) {
    // add to the table
    table.addLink(data);
  } else {
    // TODO: display error
  }
}

function init() {
  document.querySelector("#input-search").addEventListener("keyup", (event) => {
    table.searchFor(event.currentTarget.value);
  });

  document
    .querySelector("#save-new-link")
    .addEventListener("click", (event) => {
      event.preventDefault();
      saveLink();
    });

  // check access
  try {
    table._getData();
  } catch (error) {
    console.log(error);
  }
}

init();
