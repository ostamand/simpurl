export default class LinksTable {
  constructor(containerSelector) {
    this.container = document.querySelector(containerSelector);
    this.overlay = document.querySelector("#overlay-details")

    this.headers = ["URL", "Description", "Symbol"];
    this.url = "http://localhost:8001";

    this.data = null;
    this.selectedRow = null;
    this.table = null;
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
    document.querySelector("#table-links tr.table-light")?.classList.remove("table-light")
  }

  closeOverlay() {
    this.overlay.classList.remove("overlay-open")
    this.removeRowHighlight()
  }

  openOverlay() {
    this.overlay.classList.add("overlay-open")
  }

  openHover() {
    if (!this.selectedRow) {return}
    window.open(this.selectedRow.firstChild.textContent, "_blank")
  }

  /**
   * Render a new table
   */
  render() {
    this.container.innerHTML = "";

    const table = document.createElement("table");
    table.classList.add("table", "table-dark", "table-hover");
    table.id = "table-links"

    // create table header
    let header = "<thead><tr>";
    this.headers.forEach((h) => {
      header += `<th>${h}</th>`;
    });
    table.insertAdjacentHTML("beforeend", header);

    const body = document.createElement("tbody")

    // create each row
    this.data.links.forEach((link) => {
      const row = document.createElement("tr");

      // manage row clicking
      row.addEventListener("click", (event) => {
        // check if same row already selected
        if(event.currentTarget.classList.contains("table-light")) {
          this.closeOverlay()
          return
        }
        this.removeRowHighlight()
        event.currentTarget.classList.add("table-light")
        this.openOverlay()
      })

      row.addEventListener("mouseenter", event => {
        this.selectedRow = event.currentTarget
      })

      row.addEventListener("mouseleave", event => {
        this.selectedRow = null
      })

      let content = "";
      this.headers.forEach((header) => {
        content += `<td>${link[header]}</td>`;
      });

      row.innerHTML = content;

      body.appendChild(row);
    });

    table.appendChild(body)
    this.container.appendChild(table);
    this.table = table
  }

  /**
   * Update table render based on active filters
   */
  update() {}

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
