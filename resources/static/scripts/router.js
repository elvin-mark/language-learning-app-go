// Handle hamburger toggle
const toggle = document.getElementById("menuToggle");
const menu = document.getElementById("menu");

toggle.addEventListener("click", () => {
  menu.classList.toggle("open");
  toggle.classList.toggle("open");

  const expanded = toggle.getAttribute("aria-expanded") === "true";
  toggle.setAttribute("aria-expanded", !expanded);
  menu.setAttribute("aria-hidden", expanded);
});

// Handle section switching
document.querySelectorAll(".menu-list a").forEach((item) => {
  item.addEventListener("click", async () => {
    const target = item.dataset.target;

    const response = await fetch(`pages/${target}.html`);
    const html = await response.text();

    // Show selected section
    document.getElementById("content").innerHTML = html;
    if (target == "lessons") {
      lessonsEngine.fetch();
    } else if (target == "vocab") {
      vocabEngine.fetch();
    }
    if (window.innerWidth < 700) {
      // Close menu on mobile
      menu.classList.remove("open");
      toggle.classList.remove("open");
    }
  });
});
