class VocabEngine {
  page = 1;
  lastSearchTerm = "";
  pageSize = 10;
  items = [];
  constructor() {}

  async fetch(page = 1, lang = "Korean") {
    const data = await api.getVocabulary({
      userId: localStorage.getItem("userId"),
      language: lang,
      page: page,
      pageSize: this.pageSize,
    });
    this.render(data);
    this.renderContent(0);
  }

  async search(pattern, page = 1, lang = "Korean") {
    if (!pattern) {
      lastSearchTerm = "";
      return this.fetch(page, lang);
    }
    const data = await api.searchVocabulary({
      userId: localStorage.getItem("userId"),
      language: lang,
      grammarPattern: pattern,
      page: page,
      pageSize: this.pageSize,
    });
    this.render(data);
    this.renderContent(0);
  }

  async renderContent(idx) {
    const container = document.getElementById("vocabContent");
    container.innerHTML = "";

    if (this.items.length < 1) {
      return;
    }

    let item = this.items[idx];
    if (!item) {
      return;
    }

    container.innerHTML = `
    <h1>${item.Word}</h1>
    `;
  }

  generateCard(idx, item) {
    return `
    <div class="card">
        <div class="card-body">
        <h3 class="card-title">${item.Language}</h3>
        <p class="card-text">${item.Word}</p>
        <button class="card-btn" onclick="vocabEngine.renderContent(${idx})">Learn More</button>
        </div>
    </div>
`;
  }

  async render(items) {
    this.items = items;
    const container = document.getElementById("vocabList");
    container.innerHTML = "";

    if (!items || items.length === 0) {
      container.innerHTML = "<p>No results found.</p>";
      return;
    }

    let list = "";
    items.forEach((g, idx) => {
      list += this.generateCard(idx, g);
    });
    container.innerHTML = list;
  }

  nextPage() {
    this.page++;
    document.getElementById("vocabCurrentPage").textContent = this.page;
    let lang = document.getElementById("vocabLanguage").value;

    if (lastSearchTerm) {
      this.search(lastSearchTerm, this.page, lang);
    } else {
      this.fetch(this.page, lang);
    }
  }

  prevPage() {
    if (this.page > 1) this.page--;
    document.getElementById("vocabCurrentPage").textContent = this.page;
    let lang = document.getElementById("vocabLanguage").value;

    if (lastSearchTerm) {
      this.search(lastSearchTerm, this.page, lang);
    } else {
      this.fetch(this.page, lang);
    }
  }

  onClickSearch() {
    const value = document.getElementById("vocabSearch").value.trim();
    this.page = 1;
    document.getElementById("vocabCurrentPage").textContent = 1;
    let lang = document.getElementById("vocabLanguage").value;
    this.search(value, this.page, lang);
  }
}

var vocabEngine = new VocabEngine();
