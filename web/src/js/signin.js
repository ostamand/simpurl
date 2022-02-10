import getConfigs from "./defaults.js";
import AnimateTitle from "./common/animate-title.js";
import AnimateDots from "./common/animate-dots.js";

const usernameInput = document.querySelector("#input-username");
const passwordInput = document.querySelector("#input-password");
const main = document.querySelector("main");

const configs = getConfigs();

const title = new AnimateTitle(
  document.querySelector("#signup-title"),
  "SimpURL"
).start();

const submitButton = new AnimateDots(
  document.querySelector("#submit-btn"),
  "Sign in"
);

document.querySelector("#form-signin").addEventListener("submit", (event) => {
  event.preventDefault();
  signin();
});

const displayAlert = (type, message) => {
  const alert = document.querySelector("#alert");
  if (alert != null) {
    alert.remove();
  }
  main.insertAdjacentHTML(
    "afterbegin",
    /*html*/
    `
      <div id="alert"
      class="alert alert-${type} alert-dismissible fade show"
      role="alert">
      ${message}
      <button
        type="button"
        class="btn-close"
        data-bs-dismiss="alert"
        aria-label="Close">
      </button>
    </div>
  `
  );
};

const init = () => {
  const alert = window.localStorage.getItem("alert");
  if (alert != null) {
    [type, message] = alert.split(";");
    displayAlert(type, message);
    window.localStorage.removeItem("alert");
  }
};

async function signin() {
  try {
    submitButton.start();

    const response = await fetch(configs.apiEndpoint + "/signin", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: usernameInput.value,
        password: passwordInput.value,
      }),
    });
    if (response.status != 200) {
      submitButton.stop();
      displayAlert("danger", "Wrong! Try again.");
      return;
    }
    // TODO: change signin so that it returns the expire date
    const data = await response.json();

    // set or update local storage
    window.localStorage.setItem("username", data.username);
    window.localStorage.setItem("email", data.email);

    title.stop().highlightLastN(3);
    setTimeout(() => {
      window.location.replace("/index.html");
    }, 1500);
  } catch (error) {
    // TODO: add alert for user
    console.log(error);
  }
}

init();
