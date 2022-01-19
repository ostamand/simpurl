import LinksTable from "./links-table.js";

const table = new LinksTable("#container-links");

document.querySelector("#input-search").addEventListener("keyup", (event) => {
  table.searchFor(event.currentTarget.value);
});

try {
  table.getData();
} catch (error) {
  console.log(error);
}
