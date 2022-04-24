import { LitElement, html } from "lit";

export default class TagsToolbar extends LitElement {
  static properties = {
    tags: {},
  };

  constructor() {
    super();
    this.tags = [];
  }

  setTags(tags) {
    this.tags = tags;
    this.requestUpdate("tags");
  }

  _add() {
    //TODO define labelElement only once
    let labelElement = document.querySelector("input.new-tag-label");
    this.tags.push(labelElement.value);
    this.requestUpdate("tags");
  }

  _deleteTag(event) {
    let labelName = event.target.parentElement.previousElementSibling.innerText;
    // remove label name from tags
    this.tags = this.tags.filter((label) => label != labelName);
    console.log(this.tags);
    this.requestUpdate("tags");
  }

  _renderTags() {
    return this.tags.map((tag) => {
      return html` <li class="nav-item px-1">
        <span class="badge rounded-pill bg-secondary">${tag}</span>
        <a
          style="cursor:pointer;"
          class="text-muted"
          @click="${this._deleteTag}"
          ><small>x</small></a
        >
      </li>`;
    });
  }

  render() {
    return html`<nav class="py-0 navbar navbar-expand-md navbar-dark bg-dark">
      <div class="container-fluid">
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNavDarkDropdown"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse">
          <ul class="navbar-nav align-items-center">
            <li class="nav-item dropdown">
              <a
                class="nav-link dropdown-toggle"
                href="#"
                id="navbarDarkDropdownMenuLink"
                role="button"
                data-bs-toggle="dropdown"
              >
                Tags
              </a>
              <ul class="dropdown-menu dropdown-menu-dark">
                <li class="p-1">
                  <input
                    type="text"
                    class="form-control new-tag-label"
                    placeholder="Name"
                  />
                </li>
                <li class="mt-1 mx-1">
                  <button class="btn btn-primary" @click="${this._add}">
                    Add
                  </button>
                </li>
              </ul>
            </li>
            ${this._renderTags()}
          </ul>
        </div>
      </div>
    </nav>`;
  }

  createRenderRoot() {
    return this;
  }
}
