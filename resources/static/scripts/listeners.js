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
        currentLessonId = lessonId;
      } else {
        console.error(`Lesson with ID ${lessonId} not found in current data.`);
        alert("Could not load lesson details.");
      }
    }
  }
});

// Event listener for the modal close button
closeLessonModalButton.addEventListener("click", closeLessonDetailModal);
closeExerciseModalButton.addEventListener("click", closeExerciseModal);

// Close modal if clicking outside of the modal content
window.addEventListener("click", (event) => {
  if (event.target === lessonDetailModal) {
    closeLessonDetailModal();
  }
  if (event.target === exerciseModal) {
    closeExerciseModal();
  }
});

// User profile settings
const saveSettings = () => {
  updateUserProfile(prefLang.value, targetLang.value).then((resp) => {
    loadDashboardData();
  });
};

prefLang.addEventListener("change", saveSettings);
targetLang.addEventListener("change", saveSettings);

generateDialogueInitBtn.addEventListener("click", async () => {
  if (!localStorage.getItem(AUTH_TOKEN_KEY)) {
    alert("You need to be logged in to generate exercises.");
    return;
  }
  // Disable button to prevent multiple clicks
  generateDialogueInitBtn.disabled = true;
  generateDialogueInitBtn.textContent = "Generating...";
  try {
    const exercise = await generateDialogueInitExercise(currentLessonId);
    renderDialogueInitExercise(exercise);
    currentExercise = exercise;
    openExerciseModal();
    closeLessonDetailModal(); // Close lesson detail modal after generating exercise
  } catch (error) {
    console.error("Error generating Dialogue Init Exercise:", error);
    alert("Failed to generate Dialogue Init Exercise. Please try again.");
  } finally {
    generateDialogueInitBtn.disabled = false;
    generateDialogueInitBtn.textContent = "Dialogue Init";
  }
});

submitExerciseBtn.addEventListener("click", async () => {
  if (!localStorage.getItem(AUTH_TOKEN_KEY)) {
    alert("You need to be logged in to generate exercises.");
    return;
  }
  // Disable button to prevent multiple clicks
  submitExerciseBtn.disabled = true;
  submitExerciseBtn.textContent = "Generating...";
  try {
    const history = ""; // TODO: Implement a way to get history from user
    const exercise = await generateDialogueContinuationExercise(
      currentLessonId,
      history
    );
    renderDialogueContinuationExercise(exercise);
    openExerciseModal();
    closeLessonDetailModal(); // Close lesson detail modal after generating exercise
  } catch (error) {
    console.error("Error generating Dialogue Continue Exercise:", error);
    alert("Failed to generate Dialogue Continue Exercise. Please try again.");
  } finally {
    submitExerciseBtn.disabled = false;
    submitExerciseBtn.textContent = "Dialogue Continue";
  }
});

generateReadingComprehensionBtn.addEventListener("click", async () => {
  if (!localStorage.getItem(AUTH_TOKEN_KEY)) {
    alert("You need to be logged in to generate exercises.");
    return;
  }
  // Disable button to prevent multiple clicks
  generateReadingComprehensionBtn.disabled = true;
  generateReadingComprehensionBtn.textContent = "Generating...";
  try {
    const exercise = await generateReadingComprehensionExercise(
      currentLessonId
    );
    renderReadingComprehensionExercise(exercise);
    currentExercise = exercise;
    openExerciseModal();
    closeLessonDetailModal(); // Close lesson detail modal after generating exercise
  } catch (error) {
    console.error("Error generating Reading Comprehension Exercise:", error);
    alert(
      "Failed to generate Reading Comprehension Exercise. Please try again."
    );
  } finally {
    generateReadingComprehensionBtn.disabled = false;
    generateReadingComprehensionBtn.textContent = "Reading Comprehension";
  }
});

function gradeReadingComprehensionResponse(idx) {
  let inputs = document.getElementsByClassName("question-answer-input");
  console.log(currentExercise.questions[idx]);
  console.log(inputs[idx].value);
  console.log(currentExercise.short_text);
}

generateTranslationBtn.addEventListener("click", async () => {
  if (!localStorage.getItem(AUTH_TOKEN_KEY)) {
    alert("You need to be logged in to generate exercises.");
    return;
  }
  // Disable button to prevent multiple clicks
  generateTranslationBtn.disabled = true;
  generateTranslationBtn.textContent = "Generating...";
  try {
    const exercise = await generateTranslationExercise(currentLessonId);
    currentExercise = exercise;
    renderTranslationExercise(exercise); // Pass lessonId for grading
    openExerciseModal();
    closeLessonDetailModal(); // Close lesson detail modal after generating exercise
  } catch (error) {
    console.error("Error generating Translation Exercise:", error);
    alert("Failed to generate Translation Exercise. Please try again.");
  } finally {
    generateTranslationBtn.disabled = false;
    generateTranslationBtn.textContent = "Translation";
  }
});

function gradeTranslation(idx) {
  let inputs = document.getElementsByClassName("sentence-translation-input");
  gradeTranslationExercise(
    currentLessonId,
    currentExercise.sentences[idx],
    inputs[idx].value
  ).then((data) => {
    // TODO: Display the feedback and scores in the UI
    console.log(data);
  });
}

// --- Initial Load ---

document.addEventListener("DOMContentLoaded", () => {
  checkAuthStatus();
});
