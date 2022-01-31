import { LitElement, html } from "lit";
import { getSessionToken, clearSessionToken } from "../helpers";

export default class SigninElement extends LitElement {
  static properties = {
    session: {},
    username: {},
  };

  constructor() {
    super();
    this.session = getSessionToken();
    this.username = window.localStorage.getItem("username");
  }

  render() {
    if (this.session.length > 0) {
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

  _signout(event) {
    clearSessionToken();
    window.localStorage.clear("username");
    window.localStorage.clear("session");
    window.localStorage.setItem("alert", "primary;Please sign in.");
    document.location.replace("/signin.html");
  }

  createRenderRoot() {
    return this;
  }
}
