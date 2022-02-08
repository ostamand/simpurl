import { formatURL } from "../helpers.js";
import FetchWrapper from "../common/fetch-wrapper.js";

const API = new FetchWrapper();

export default class TableOverlay {
  constructor() {
    this.overlay = document.querySelector("#overlay-details");
    this.closeCallbacks = [];
    this.link = null;
    this.updateCb = null;

    this.title = document.querySelector("#overlay-title");
    this.description = document.querySelector("#overlay-description");
    this.symbol = document.querySelector("#overlay-symbol");
    this.note = document.querySelector("#overlay-note");

    document
      .querySelector("#btn-overlay-close")
      .addEventListener("click", () => {
        this.close();
      });

    document
      .querySelector("#btn-overlay-update")
      .addEventListener("click", () => {
        this.update();
      });

    document.addEventListener("keyup", (event) => {
      if (event.key == "Escape") {
        this.close();
      }
    });
  }

  display(link) {
    this.link = link;

    this.title.textContent = formatURL(link.url);
    this.title.setAttribute("href", link.url);

    this.description.value = link.description;
    this.symbol.value = link.symbol;
    // TODO: not available -> this.note.textContent =
  }

  async update() {
    const data = {
      description: this.description.value,
      symbol: this.symbol.value,
    };
    const [status, _] = await API.patch(`/urls/${this.link.urlID}`, data);
    console.log(status);

    if (status != 200) {
      // TODO: display error
      return;
    }

    // update table & overlay data
    for (const field in data) {
      this.link[field] = data[field];
    }
    this.updateCb(this.link);
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

  onUpdate(f) {
    this.updateCb = f;
  }
}
