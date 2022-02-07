const token = window.localStorage.getItem("session");
if (!token) {
  window.localStorage.setItem("alert", "danger;Please sign in.");
  window.location.replace("/signin.html");
}
