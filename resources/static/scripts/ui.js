function updateUIForLoggedInState(username) {
  loginSection.style.display = "none";
  dashboardSection.style.display = "block";
  authNav.style.display = "block"; // Ensure nav is visible
  loginBtn.style.display = "none";
  logoutBtn.style.display = "inline-block";
  currentUsernameSpan.textContent = username;
}

function updateUIForLoggedOutState() {
  dashboardSection.style.display = "none";
  loginSection.style.display = "block";
  authNav.style.display = "block"; // Ensure nav is visible
  loginBtn.style.display = "inline-block";
  logoutBtn.style.display = "none";
  currentUsernameSpan.textContent = "";
  // Clear any lingering modal content
  closeLessonDetailModal();
}

function clearDashboardData() {
  lessonsList.innerHTML = "";
  vocabularyList.innerHTML = "";
  grammarList.innerHTML = "";
  currentLessonsData = [];
  currentWordsData = [];
  currentGrammarData = [];
}

async function checkAuthStatus() {
  const token = localStorage.getItem(AUTH_TOKEN_KEY);
  const username = localStorage.getItem(USERNAME_KEY);

  if (token && username) {
    // Optionally, verify the token with a user profile request
    try {
      const user = await getUserProfile(token);
      if (user) {
        updateUIForLoggedInState(username);
        await loadDashboardData(token);
        return;
      }
    } catch (error) {
      console.error("Token validation failed, logging out.", error);
      // If token is invalid, clear it and proceed to login
      logoutUser();
    }
  }
  updateUIForLoggedOutState();
}

// --- Modal Functions ---

function openLessonDetailModal() {
  lessonDetailModal.style.display = "block";
}

function closeLessonDetailModal() {
  lessonDetailModal.style.display = "none";
}

// --- Rendering Functions ---

function renderUserProfile(user) {
  if (user && user.username) {
    currentUsernameSpan.textContent = user.username;
  }
}

function renderLessons(lessons) {
  lessonsList.innerHTML = ""; // Clear existing lessons
  currentLessonsData = lessons; // Store lessons globally

  if (lessons && lessons.length > 0) {
    lessons.forEach((lesson) => {
      const li = document.createElement("li");
      li.textContent = `Lesson ${lesson.id}: ${
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
  startExerciseBtn.dataset.lessonId = lesson.id;
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

// --- Event Listeners ---

loginForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  const username = usernameInput.value;
  const password = passwordInput.value;
  await loginUser(username, password);
});

logoutBtn.addEventListener("click", logoutUser);

generateLessonBtn.addEventListener("click", async () => {
  // Ensure user is logged in before proceeding
  if (!localStorage.getItem(AUTH_TOKEN_KEY)) {
    alert("You need to be logged in to generate lessons.");
    return;
  }

  // Disable button to prevent multiple clicks
  generateLessonBtn.disabled = true;
  generateLessonBtn.textContent = "Generating...";

  try {
    await generateNewLesson();
    // After generating, refresh the lessons list to show the new lesson
    await loadDashboardData(); // Reload all dashboard data to reflect changes
    alert("New lesson generated successfully!");
  } catch (error) {
    console.error("Error generating new lesson:", error);
    alert("Failed to generate new lesson. Please try again.");
    // Re-enable button on error
    generateLessonBtn.disabled = false;
    generateLessonBtn.textContent = "Generate New Lesson";
  } finally {
    // Re-enable button and reset text regardless of success or failure
    // Use a slight delay to ensure the UI updates after alert closes.
    setTimeout(() => {
      generateLessonBtn.disabled = false;
      generateLessonBtn.textContent = "Generate New Lesson";
    }, 100);
  }
});

// Event listener for clicks on the lessons list (using event delegation)
lessonsList.addEventListener("click", (event) => {
  if (event.target.classList.contains("lesson-item")) {
    const lessonId = parseInt(event.target.dataset.lessonId);
    if (!isNaN(lessonId)) {
      const lesson = currentLessonsData.find((l) => l.Id === lessonId);
      if (lesson) {
        populateLessonDetailModal(lesson);
        openLessonDetailModal();
      } else {
        console.error(`Lesson with ID ${lessonId} not found in current data.`);
        alert("Could not load lesson details.");
      }
    }
  }
});

// Event listener for the modal close button
closeModalButton.addEventListener("click", closeLessonDetailModal);

// Close modal if clicking outside of the modal content
window.addEventListener("click", (event) => {
  if (event.target === lessonDetailModal) {
    closeLessonDetailModal();
  }
});

// Event listener for the start exercise button
startExerciseBtn.addEventListener("click", async () => {
  const lessonId = parseInt(startExerciseBtn.dataset.lessonId);
  if (isNaN(lessonId)) {
    alert("Invalid lesson selected for exercise.");
    return;
  }

  // Placeholder for starting an exercise.
  // TODO: Implement logic to call appropriate exercise generation API based on lesson type or user choice.
  alert(
    `Starting exercise for Lesson ID: ${lessonId}. This feature is under development.`
  );
  // Example: You might call something like:
  // try {
  //     const exercise = await generateTranslationExercise(lessonId);
  //     // then navigate to an exercise page or display exercise here
  // } catch (error) {
  //     alert('Failed to start exercise.');
  // }

  closeLessonDetailModal(); // Close modal after initiating exercise
});

// --- Initial Load ---

document.addEventListener("DOMContentLoaded", () => {
  checkAuthStatus();
});
