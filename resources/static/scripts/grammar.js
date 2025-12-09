let grammarPage = 1;
let lastSearchTerm = "";

class GrammarEngine {
  constructor() {}

  async fetch(page = 1) {
    const data = await api.getGrammar({
      userId: 1,
      language: "Korean",
      page: page,
      pageSize: 10,
    });
    this.render(data);
  }

  async search(pattern, page = 1) {
    if (!pattern) {
      lastSearchTerm = "";
      return this.fetch(page);
    }
    const data = await api.searchGrammar({
      userId: 1,
      language: "Korean",
      pattern: pattern,
      page: page,
      pageSize: 10,
    });
    this.render(data);
  }

  async render(items) {
    const container = document.getElementById("grammarList");
    container.innerHTML = "";

    if (!items || items.length === 0) {
      container.innerHTML = "<p>No results found.</p>";
      return;
    }

    items.forEach((g) => {
      const div = document.createElement("div");
      div.className = "grammar-item";
      div.innerHTML = `
      <strong>${g.Pattern}</strong><br>
      Mastery Score: ${g.MasteryScore}<br>
      Last Reviewed: ${g.LastReviewed}
    `;
      container.appendChild(div);
    });
  }

  nextPage() {
    grammarPage++;
    document.getElementById("currentPage").textContent = grammarPage;

    if (lastSearchTerm) {
      this.search(lastSearchTerm, grammarPage);
    } else {
      this.fetch(grammarPage);
    }
  }

  prevPage() {
    if (grammarPage > 1) grammarPage--;
    document.getElementById("currentPage").textContent = grammarPage;

    if (lastSearchTerm) {
      this.search(lastSearchTerm, grammarPage);
    } else {
      this.grammar(grammarPage);
    }
  }

  onClickSearch() {
    const value = document.getElementById("grammarSearch").value.trim();
    grammarPage = 1;
    document.getElementById("currentPage").textContent = 1;
    this.search(value, grammarPage);
  }
}

var grammarEngine = new GrammarEngine();
