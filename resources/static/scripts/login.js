async function login() {
  var username = document.getElementById("username").value;
  var password = document.getElementById("password").value;

  const res = await fetch("/auth/token", {
    method: "POST",
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    document.getElementById("username").value = "";
    document.getElementById("password").value = "";
  } else {
    res.json().then((data) => {
      localStorage.setItem("token", data.AccessToken);
      localStorage.setItem("userId", data.UserId);
      window.location.href = "/";
    });
  }
}
