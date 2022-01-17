import LinksTable from "./links-table.js";

const table = new LinksTable("#container-links");

document.querySelector("#btn-overlay-close").addEventListener("click", () => {
  table.closeOverlay();
});

document.querySelector("#input-search").addEventListener("keyup", (event) => {
  table.searchFor(event.currentTarget.value);
});

document.addEventListener("keyup", (event) => {
  if (event.key == "o") {
    table.openLinkHover();
  } else if (event.key == "Escape") {
    table.closeOverlay();
  }
});

try {
  table.getData();
} catch (error) {
  console.log(error);
}
