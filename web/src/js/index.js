import TableURL from "./url/table.js";
import SigninElement from "./component/signin-toolbar.js";
import TagsToolbar from "./component/tags-toolbar.js";
import FetchWrapper from "./common/fetch-wrapper.js";

customElements.define("signin-btn", SigninElement);
customElements.define("tags-toolbar", TagsToolbar);

const table = new TableURL("#container-links");
const API = new FetchWrapper();

async function saveLink() {
  const request = {};
  document
    .querySelectorAll("#form-new-link input, #form-new-link textarea")
    .forEach((input) => {
      request[input.name] = input.value;
    });

  //TODO remove and add proper way of tagging new URL

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
