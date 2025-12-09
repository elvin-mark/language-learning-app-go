class LanguageLearningAPI {
  constructor(baseURL = "http://localhost:8081") {
    this.baseURL = baseURL;
  }

  // Helper for GET with query params
  async get(path, params = {}) {
    const url = new URL(this.baseURL + path);
    Object.keys(params).forEach((k) => {
      if (params[k] !== undefined && params[k] !== null) {
        url.searchParams.append(k, params[k]);
      }
    });

    const token = localStorage.getItem("token");

    const res = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Basic " + token,
      },
    });

    if (!res.ok) throw new Error(await res.text());
    return res.json();
  }

  // Helper for POST with JSON body
  async post(path, bodyObj) {
    const token = localStorage.getItem("token");

    const res = await fetch(this.baseURL + path, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Basic " + token,
      },
      body: JSON.stringify(bodyObj),
    });

    if (!res.ok) throw new Error(await res.text());
    return res.json();
  }

  // ---- Agents ----

  getExercises(data) {
    // { lang, lessonId, practicePattern }
    return this.post("/agents/exercises", data);
  }

  getLesson(data) {
    // { lang, userId }
    return this.post("/agents/lessons", data);
  }

  // ---- Grammar ----

  getGrammar({ userId, language, page, pageSize }) {
    return this.get("/resources/grammar", {
      userId,
      language,
      page,
      pageSize,
    });
  }

  searchGrammar({ userId, language, pattern, page, pageSize }) {
    return this.get("/resources/grammar/search", {
      userId,
      language,
      pattern,
      page,
      pageSize,
    });
  }

  // ---- Lessons ----

  getLessons({ userId, language, page, pageSize }) {
    return this.get("/resources/lessons", {
      userId,
      language,
      page,
      pageSize,
    });
  }

  searchLessons({ userId, language, grammarPattern, page, pageSize }) {
    return this.get("/resources/lessons/search", {
      userId,
      language,
      grammarPattern,
      page,
      pageSize,
    });
  }

  // ---- Vocabulary ----

  getVocabulary({ userId, language, page, pageSize }) {
    return this.get("/resources/vocabulary", {
      userId,
      language,
      page,
      pageSize,
    });
  }
}

var api = new LanguageLearningAPI();
