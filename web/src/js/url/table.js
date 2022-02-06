import TableOverlay from "./table-overlay.js";
import { formatURL } from "../helpers.js";
import configs from "../defaults.js"

export default class TableURL {
  constructor(containerSelector) {
    this.container = document.querySelector(containerSelector);

    this.headers = ["url", "description", "symbol"];
    this.formatHeaders = {
      URL: formatURL,
    };

    this.url = configs.apiEndpoint; // TODO get from configs

    this.data = null;
    this.selectedRow = null;
    this.table = null;
    this.filters = {};

    this.overlay = new TableOverlay();
    this.overlay.onClose(() => this.removeRowHighlight());

    document.addEventListener("keyup", (event) => {
      if (event.key == "o") {
        this.openLinkHover();
      }
    });
  }

  removeRowHighlight() {
    document
      .querySelector("#table-links tr.table-light")
      ?.classList.remove("table-light");
  }

  searchFor(value) {
    this.filters.search_for = value;
    this.update();
  }

  /**
   * Open the link currently being hovered
   */
  openLinkHover() {
    if (!this.selectedRow) {
      return;
    }
    window.open(this.selectedRow.firstChild.textContent, "_blank");
  }

  addLink(link) {
    this.data.links.push(link);
    // TODO: add check if table defined
    const body = document.querySelector("#table-links tbody");
    const row = this._createRow(link);

    row.classList.add("table-success");
    setTimeout(() => {
      row.classList.remove("table-success");
    }, 2000);

    body.prepend(row);

    // close modal
    document.querySelector("#close-new-link").click();
  }

  _createRow(url) {
    const row = document.createElement("tr");

    // manage row clicking
    row.addEventListener("click", (event) => {
      const target = event.currentTarget;
      // check if same row already selected
      if (target.classList.contains("table-light")) {
        this.removeRowHighlight();
        this.overlay.close();
        return;
      }
      this.removeRowHighlight();
      target.classList.add("table-light");

      // find what to display by id
      const url = this.data.find(
        (x) => x.urlID === Number.parseInt(target.dataset.id)
      );
      this.overlay.display(url);
      this.overlay.open();
    });

    row.addEventListener("mouseenter", (event) => {
      this.selectedRow = event.currentTarget;
    });

    row.addEventListener("mouseleave", () => {
      this.selectedRow = null;
    });

    let content = "";
    this.headers.forEach((header) => {
      row.setAttribute(`data-${header}`, url[header]); // used by the search

      // apply formatting if necessary
      let text = url[header];
      if (header in this.formatHeaders) {
        text = this.formatHeaders[header](text);
      }
      content += `<td>${text}</td>`;
    });
    row.setAttribute("data-id", url.urlID); // id used when showing the overlay

    row.innerHTML = content;

    return row;
  }

  /**
   * Render a new table
   */
  render() {
    this.container.innerHTML = "";

    const table = document.createElement("table");
    table.classList.add("table", "table-dark", "table-hover", "align-middle");
    table.id = "table-links";

    // create table header
    let header = "<thead><tr>";
    this.headers.forEach((h) => {
      header += `<th>${h}</th>`;
    });
    table.insertAdjacentHTML("beforeend", header);

    const body = document.createElement("tbody");

    // create each row
    this.data.forEach((link) => {
      const row = this._createRow(link);
      body.appendChild(row);
    });

    table.appendChild(body);
    this.container.appendChild(table);
    this.table = table;
  }

  /**
   * Update table render based on active filters
   */
  update() {
    this.table.lastChild.childNodes.forEach((row) => {
      const dataset = row.dataset;
      let cond = true;

      if ("search_for" in this.filters) {
        let searchContent = "";
        Object.keys(dataset).forEach((key) => {
          searchContent += " " + dataset[key];
        });
        cond = cond && searchContent.includes(this.filters["search_for"]);
      }

      // set visibility of row based on final cond value
      if (cond) {
        row.classList.remove("hide");
      } else {
        row.classList.add("hide");
      }
    });
  }

  /**
   * Get data from API
   * GET /api/links
   */
  async _getData() {
    const limit = -1;
    const response = await fetch(this.url + "/urls", {
      method: "GET",
      credentials: "include",
    });
    if (response.status != 200) {
      // assuming session token is no good anymore
      // TODO: fix this because can be something else
      window.location.replace("/signin.html");
      return;
    }
    this.data = await response.json();
    this.render();
  }
}
