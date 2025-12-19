// --- Rendering Functions ---

function renderUserProfile(user) {
  if (user && user.Username) {
    currentUsernameSpan.textContent = user.Username;
    prefLang.value = user.PreferredLanguage;
    targetLang.value = user.TargetLanguage;
  }
}

function renderUserStatusReport(data) {
  if (data) {
    const wordPercent = (data.MasteredWords / data.TotalWords) * 100;
    const grammarPercent =
      (data.MasteredGrammarPatterns / data.TotalGrammarPatterns) * 100;

    const wordCount = `${data.MasteredWords} / ${data.TotalWords}`;
    const grammarCount = `${data.MasteredGrammarPatterns} / ${data.TotalGrammarPatterns}`;

    document.getElementById("word-fill").style.width = `${wordPercent}%`;
    document.getElementById("grammar-fill").style.width = `${grammarPercent}%`;
    document.getElementById("word-count").innerHTML = wordCount;
    document.getElementById("grammar-count").innerHTML = grammarCount;
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

  try {
    lessonDetailContent.innerHTML = marked.parse(lesson.Content);
  } catch {
    lessonDetailContent.textContent = lesson.Content;
  }
  lessonDetailGrammar.textContent = lesson.Grammar;

  try {
    let sampleSentences = JSON.parse(lesson.SampleSentences);
    lessonDetailSampleSentences.innerHTML = sampleSentences
      .map((elem) => "<div>" + elem + "</div>")
      .join("");
  } catch {
    lessonDetailSampleSentences.textContent = lesson.SampleSentences || "N/A";
  }

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

  try {
    let wordsMeaning = JSON.parse(lesson.WordsMeaning);
    let wordsMeaningContent = "";
    for (word in wordsMeaning) {
      wordsMeaningContent += `<div>
        <span>${word}: </span> ${wordsMeaning[word]}
      </div>
      `;
    }
    lessonDetailWordsMeaning.innerHTML = wordsMeaningContent;
  } catch {
    lessonDetailWordsMeaning.textContent = lesson.WordsMeaning || "N/A";
  }
}

function renderDialogueInitExercise(exercise) {
  exerciseTitle.textContent = "Dialogue Initialization Exercise";
  exerciseContent.innerHTML = `
    <p><strong>Situation:</strong> ${exercise.situation}</p>
    <p><strong>Start the dialogue:</strong></p>
    <p>${exercise.init}</p>
  `;
  exerciseInput.style.display = "block"; // Ensure single input is visible
  exerciseAnswer.placeholder = "Your response...";
  submitExerciseBtn.dataset.exerciseType = "dialogueInit"; // Store exercise type for grading
  submitExerciseBtn.style.display = "block"; // Ensure submit button is visible
  // Reset feedback display
  exerciseFeedback.style.display = "none";
  feedbackText.innerHTML = "";
  feedbackScore.innerHTML = "";
}

function renderDialogueContinuationExercise(exercise) {
  exerciseTitle.textContent = "Dialogue Continuation Exercise";
  exerciseContent.innerHTML = `
    <p><strong>Continue the dialogue:</strong></p>
    <p>${exercise.next}</p>
  `;
  exerciseInput.style.display = "block"; // Ensure single input is visible
  exerciseAnswer.placeholder = "Your response...";
  submitExerciseBtn.dataset.exerciseType = "dialogueContinue"; // Store exercise type for grading
  submitExerciseBtn.style.display = "block"; // Ensure submit button is visible
  // Reset feedback display
  exerciseFeedback.style.display = "none";
  feedbackText.innerHTML = "";
  feedbackScore.innerHTML = "";
}

function renderReadingComprehensionExercise(exercise) {
  exerciseTitle.textContent = "Reading Comprehension Exercise";
  exerciseInput.style.display = "none"; // Hide single input area
  submitExerciseBtn.style.display = "none"; // Hide single submit button
  exerciseFeedback.style.display = "none"; // Hide general feedback area

  let questionsHtml = `
    <p><strong>Read the following text and answer the questions:</strong></p>
    <p>${exercise.short_text}</p>
    <p><strong>Questions:</strong></p>
  `;

  exercise.questions.forEach((q, index) => {
    questionsHtml += `
      <div class="individual-question" data-question-index="${index}">
        <p>${index + 1}. ${q}</p>
        <input type="text" class="question-answer-input" placeholder="Your answer for question ${
          index + 1
        }..." />
        <button class="check-answer-btn" data-question="${q}" data-index="${index}" onclick="gradeReadingComprehension(${index})">Check</button>
        <div class="result-area" id="result-q-${index}"></div>
      </div>
    `;
  });
  exerciseContent.innerHTML = questionsHtml;
}

function renderTranslationExercise(exercise) {
  exerciseTitle.textContent = "Translation Exercise";
  exerciseInput.style.display = "none"; // Hide single input area
  submitExerciseBtn.style.display = "none"; // Hide single submit button
  exerciseFeedback.style.display = "none"; // Hide general feedback area

  let sentencesHtml = `
    <p><strong>Translate the following sentences:</strong></p>
  `;

  exercise.sentences.forEach((s, index) => {
    sentencesHtml += `
      <div class="individual-sentence" data-sentence-index="${index}">
        <p>${s}</p>
        <input type="text" class="sentence-translation-input" placeholder="Your translation for sentence ${
          index + 1
        }..." />
        <button class="check-translation-btn" data-sentence="${s}" data-index="${index}" onclick="gradeTranslation(${index})">Check</button>
        <div class="result-area" id="result-s-${index}"></div>
      </div>
    `;
  });
  exerciseContent.innerHTML = sentencesHtml;
}

function displayExerciseFeedback(grade) {
  exerciseFeedback.style.display = "block";
  feedbackText.innerHTML = `<strong>Feedback:</strong> ${grade.feedback}`;
  feedbackScore.innerHTML = `<strong>Score:</strong> ${grade.score}`;
}

async function loadDashboardData() {
  try {
    // Fetch User Profile
    const user = await getUserProfile();
    renderUserProfile(user);

    // Fetch User Status Report
    const report = await getUserStatusReport();
    renderUserStatusReport(report);

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
