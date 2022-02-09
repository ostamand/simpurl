import { LitElement, html } from "lit";
import { isLoggedIn } from "../helpers";
import FetchWrapper from "../common/fetch-wrapper";

const API = new FetchWrapper();

export default class SigninElement extends LitElement {
  static properties = {
    username: {},
  };

  constructor() {
    super();
    this.username = window.localStorage.getItem("username");
  }

  render() {
    if (isLoggedIn()) {
      return html`
        <div class="dropdown">
          <button
            class="btn btn-secondary dropdown-toggle"
            type="button"
            data-bs-toggle="dropdown"
          >
            ${this.username}
          </button>
          <ul class="dropdown-menu">
            <li>
              <a @click="${this._signout}" class="dropdown-item">Sign out</a>
            </li>
          </ul>
        </div>
      `;
    }
    return html`
      <a class="btn btn-outline-light" href="/signin.html">Sign in</a>
    `;
  }

  async _signout(event) {
    const [status, _] = await API.get("/signout");
    if (status === 200) {
      window.localStorage.clear("username");
      window.localStorage.setItem("alert", "primary;Please sign in.");
      document.location.replace("/signin.html");
    } else {
      //TODO manage signout errors.
    }
  }

  createRenderRoot() {
    return this;
  }
}
