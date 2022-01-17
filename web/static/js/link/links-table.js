export default class LinksTable {
  constructor(containerSelector) {
    this.container = document.querySelector(containerSelector);
    this.overlay = document.querySelector("#overlay-details");

    this.headers = ["URL", "Description", "Symbol"];

    this.url = "http://localhost:8001";

    this.data = null;
    this.selectedRow = null;
    this.table = null;
    this.filters = {};
  }

  _getSessionToken() {
    let session = "";
    document.cookie.split(";").forEach((cookie) => {
      const [key, value] = cookie.trim().split("=");
      if (key === "session_token") {
        session = value;
      }
    });
    return session;
  }

  removeRowHighlight() {
    document
      .querySelector("#table-links tr.table-light")
      ?.classList.remove("table-light");
  }

  closeOverlay() {
    if (this.overlay.classList.contains("overlay-open")) {
      this.overlay.classList.remove("overlay-open");
      this.removeRowHighlight();
    }
  }

  openOverlay() {
    this.overlay.classList.add("overlay-open");
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
    this.data.links.forEach((link) => {
      const row = document.createElement("tr");

      // manage row clicking
      row.addEventListener("click", (event) => {
        // check if same row already selected
        if (event.currentTarget.classList.contains("table-light")) {
          this.closeOverlay();
          return;
        }
        this.removeRowHighlight();
        event.currentTarget.classList.add("table-light");
        this.openOverlay();
      });

      row.addEventListener("mouseenter", () => {
        this.selectedRow = event.currentTarget;
      });

      row.addEventListener("mouseleave", () => {
        this.selectedRow = null;
      });

      let content = "";
      this.headers.forEach((header) => {
        row.setAttribute(`data-${header}`, link[header]);
        content += `<td>${link[header]}</td>`;
      });

      row.innerHTML = content;

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
  async getData() {
    const token = this._getSessionToken();
    const limit = -1; // TODO
    const response = await fetch(this.url + "/api/links", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({ token, limit }),
    });
    this.data = await response.json();
    this.render();
  }
}
