import LinksTable from "./links-table.js";

const table = new LinksTable("#container-links");

document.querySelector("#btn-overlay-close").addEventListener("click", () => {
  table.closeOverlay()
})

document.querySelector("#input-search").addEventListener("keyup", (event) => {
  table.searchFor(event.currentTarget.value)
})

document.addEventListener("keyup", event => {
  if(event.key == "o") {
    table.openHover()
  }
})

try {
  table.getData();
} catch (error) {
  console.log(error);
}
