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
  closeExerciseModal(); // Also close exercise modal if open
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
        await loadDashboardData();
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

function openExerciseModal() {
  if (currentExercise) {
    exerciseModal.style.display = "block";
  }
}

function closeExerciseModal() {
  exerciseModal.style.display = "none";
}

function cleanExerciseModal() {
  exerciseTitle.innerHTML = "";
  exerciseContent.innerHTML = "";
  exerciseInput.style.display = "none";
  exerciseAnswer.value = "";
  exerciseFeedback.style.display = "none";
  feedbackText.innerHTML = "";
  feedbackScore.innerHTML = "";
}
