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

// User profile settings
const saveSettings = () => {
  updateUserProfile(prefLang.value, targetLang.value).then((resp) => {
    loadDashboardData();
  });
};

prefLang.addEventListener("change", saveSettings);
targetLang.addEventListener("change", saveSettings);

// --- Initial Load ---

document.addEventListener("DOMContentLoaded", () => {
  checkAuthStatus();
});
