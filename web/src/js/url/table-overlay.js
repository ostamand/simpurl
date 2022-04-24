import { formatURL } from "../helpers.js";
import FetchWrapper from "../common/fetch-wrapper.js";

const API = new FetchWrapper();

export default class TableOverlay {
  constructor() {
    this.overlay = document.querySelector("#overlay-details");
    this.tagsToolbar = document.querySelector("#overlay-details tags-toolbar");

    this.closeCallbacks = []; //? rename this to be consistent with other callbacks?
    this.link = null;

    this.updateCb = null;
    this.deleteCb = null;

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
      .querySelector("#btn-overlay-delete")
      .addEventListener("click", () => {
        this.delete();
      });

    document
      .querySelector("#btn-overlay-undo")
      .addEventListener("click", () => {
        this.display(this.link);
      });

    document.addEventListener("keyup", (event) => {
      if (event.key == "Escape") {
        this.close();
      }
    });
  }

  async display(link) {
    this.link = link;

    this.title.textContent = formatURL(link.url);
    this.title.setAttribute("href", link.url);

    this.description.value = link.description;
    this.symbol.value = link.symbol;

    this.tagsToolbar.setTags(link.tags);

    // all the other fields could be gotten from the data
    const [status, url] = await API.get(`/urls/${this.link.urlID}`);
    if (status === 200) {
      this.note.value = url.note;
    }
  }

  getCurrentValues() {
    return {
      description: this.description.value,
      symbol: this.symbol.value,
      note: this.note.value,
    };
  }

  async updateWith(values) {
    const [status, _] = await API.patch(`/urls/${this.link.urlID}`, values);
    if (status != 200) {
      // TODO: display error
      return;
    }
    // update table & overlay data
    for (const field in values) {
      this.link[field] = values[field];
    }
    this.updateCb(this.link);
  }

  async delete() {
    // delete from the db
    // remove from the table
    // close overlay
    const [status, _] = await API.delete(`/urls/${this.link.urlID}`);
    this.deleteCb();
    this.close();
  }

  /**
   * Automatically save when closing tab (if needed)
   */
  async close() {
    const currentValues = this.getCurrentValues();

    let original = true;
    for (const property in currentValues) {
      original = original && currentValues[property] === this.link[property];
    }

    // check if tags changed also
    this.tagsToolbar.tags.forEach((tag) => {
      original = original && this.link.tags.includes(tag);
    });
    this.link.tags.forEach((tag) => {
      original = original && this.tagsToolbar.tags.includes(tag);
    });

    if (!original) {
      console.log("update");
      currentValues["tags"] = this.tagsToolbar.tags;
      await this.updateWith(currentValues);
    }
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

  onDelete(f) {
    this.deleteCb = f;
  }
}
