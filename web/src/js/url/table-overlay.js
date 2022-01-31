import { formatURL } from "../helpers.js";

export default class TableOverlay {
  constructor() {
    this.overlay = document.querySelector("#overlay-details");
    this.closeCallbacks = [];
    this.link = null;

    this.title = document.querySelector("#overlay-title");
    this.description = document.querySelector("#overlay-description");
    this.symbol = document.querySelector("#overlay-symbol");
    this.note = document.querySelector("#overlay-note");

    document
      .querySelector("#btn-overlay-close")
      .addEventListener("click", () => {
        this.close();
      });

    document.addEventListener("keyup", (event) => {
      if (event.key == "Escape") {
        this.close();
      }
    });
  }

  display(link) {
    this.link = link;

    this.title.textContent = formatURL(link.URL);
    this.title.setAttribute("href", link.URL);

    this.description.value = link.Description;

    if (link.Symbol) {
      this.symbol.value = `simpurl/${link.Symbol}`;
    }

    // TODO: not available -> this.note.textContent =
  }

  close() {
    this.overlay.classList.remove("start-50");
    this.closeCallbacks.forEach((f) => f());
  }

  open() {
    this.overlay.classList.add("start-50");
  }

  onClose(f) {
    this.closeCallbacks.push(f);
  }
}
