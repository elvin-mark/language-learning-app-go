let vocabPage = 1;
const vocabPageSize = 10;

class VocabularyEngine {
  constructor() {}

  async fetch(page = 1) {
    const data = await api.getVocabulary({
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
    const data = await api.getVocabulary({
      userId: 1,
      language: "Korean",
      page: page,
      pageSize: 10,
    });
    this.render(data);
  }

  render(items) {
    const container = document.getElementById("vocabList");
    container.innerHTML = "";

    if (!items.length) {
      container.innerHTML = "<p>No vocabulary found.</p>";
      return;
    }

    items.forEach((v) => {
      const div = document.createElement("div");
      div.className = "vocab-item";
      div.innerHTML = `
        <h3>${v.Word}</h3>
        <p><strong>Language:</strong> ${v.Language.toUpperCase()}</p>
      `;
      container.appendChild(div);
    });
  }

  nextPage() {
    vocabPage++;
    document.getElementById("vocabularyCurrentPage").textContent = vocabPage;

    if (lastSearchTerm) {
      this.search(lastSearchTerm, vocabPage);
    } else {
      this.fetch(vocabPage);
    }
  }

  prevPage() {
    if (vocabPage > 1) vocabPage--;
    document.getElementById("vocabularyCurrentPage").textContent = vocabPage;

    if (lastSearchTerm) {
      this.search(lastSearchTerm, vocabPage);
    } else {
      this.grammar(vocabPage);
    }
  }

  onClickSearch() {
    const value = document.getElementById("vocabularySearch").value.trim();
    vocabPage = 1;
    document.getElementById("vocabularyCurrentPage").textContent = 1;
    this.search(value, vocabPage);
  }
}

var vocabularyEngine = new VocabularyEngine();
