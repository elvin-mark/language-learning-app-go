// --- Rendering Functions ---

function renderUserProfile(user) {
  if (user && user.Username) {
    currentUsernameSpan.textContent = user.Username;
    prefLang.value = user.PreferredLanguage;
    targetLang.value = user.TargetLanguage;
  }
}

function renderLessons(lessons) {
  lessonsList.innerHTML = ""; // Clear existing lessons
  currentLessonsData = lessons; // Store lessons globally

  if (lessons && lessons.length > 0) {
    lessons.forEach((lesson) => {
      const li = document.createElement("li");
      li.textContent = `Lesson ${lesson.Id}: ${
        lesson.Language
      } - ${lesson.Content.substring(0, 60)}...`; // Display snippet
      li.dataset.lessonId = lesson.Id; // Add lesson ID as data attribute
      li.classList.add("lesson-item"); // Add class for easier event delegation
      lessonsList.appendChild(li);
    });
  } else {
    lessonsList.innerHTML = "<li>No lessons available yet.</li>";
  }
}

function renderVocabulary(words) {
  vocabularyList.innerHTML = ""; // Clear existing words
  currentWordsData = words; // Store words globally
  if (words && words.length > 0) {
    words.forEach((word) => {
      const li = document.createElement("li");
      // Assuming word object has 'word' and 'score' properties
      li.textContent = `${word.Word} (Score: ${word.Score})`;
      vocabularyList.appendChild(li);
    });
  } else {
    vocabularyList.innerHTML = "<li>No vocabulary learned yet.</li>";
  }
}

function renderGrammar(grammarPatterns) {
  grammarList.innerHTML = ""; // Clear existing grammar
  currentGrammarData = grammarPatterns; // Store grammar globally
  if (grammarPatterns && grammarPatterns.length > 0) {
    grammarPatterns.forEach((pattern) => {
      const li = document.createElement("li");
      // Assuming grammar object has 'pattern' and 'score' properties
      li.textContent = `${pattern.Pattern} (Score: ${pattern.Score})`;
      grammarList.appendChild(li);
    });
  } else {
    grammarList.innerHTML = "<li>No grammar patterns learned yet.</li>";
  }
}

function populateLessonDetailModal(lesson) {
  lessonDetailTitle.textContent = `Lesson ${lesson.Id} Details`;
  lessonDetailLanguage.textContent = lesson.Language;
  lessonDetailContent.textContent = lesson.Content;
  lessonDetailGrammar.textContent = lesson.Grammar;
  lessonDetailSampleSentences.textContent = lesson.SampleSentences;

  // Clear previous words
  lessonDetailWordsList.innerHTML = "";
  if (lesson.Words && lesson.Words.length > 0) {
    lesson.Words.forEach((word) => {
      const li = document.createElement("li");
      li.textContent = word;
      lessonDetailWordsList.appendChild(li);
    });
  } else {
    const li = document.createElement("li");
    li.textContent = "No specific words listed for this lesson.";
    lessonDetailWordsList.appendChild(li);
  }
  lessonDetailWordsMeaning.textContent = lesson.WordsMeaning || "N/A";

  // Store lesson ID on the start button for later use
  startExerciseBtn.dataset.lessonId = lesson.Id;
}

async function loadDashboardData() {
  try {
    // Fetch User Profile
    const user = await getUserProfile();
    renderUserProfile(user);

    // Fetch Lessons
    const lessons = await getLessons();
    renderLessons(lessons);

    // Fetch Vocabulary
    const words = await getVocabulary();
    renderVocabulary(words);

    // Fetch Grammar
    const grammarPatterns = await getGrammar();
    renderGrammar(grammarPatterns);
  } catch (error) {
    console.error("Error loading dashboard data:", error);
    // Errors like 401 will be handled by apiFetch which calls logoutUser
    // For other errors, show a general message if not already handled.
    if (
      !error.message.includes("Not authenticated") &&
      !error.message.includes("Authentication failed")
    ) {
      alert(
        "Failed to load dashboard data. Please check your connection or try logging in again."
      );
    }
  }
}
