const token = localStorage.getItem("token");
if (token == null || token == "") {
  window.location.href = "/ui/login.html";
}
