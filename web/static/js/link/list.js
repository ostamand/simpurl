import LinksTable from "./links-table.js";

const table = new LinksTable("#container-links");

document.querySelector("#btn-overlay-close").addEventListener("click", (event) => {
  table.closeOverlay()
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
