const API_BASE_URL = "";
const AUTH_TOKEN_KEY = "authToken";
const USER_ID_KEY = "userId";
const USERNAME_KEY = "username";

// DOM Elements
const loginSection = document.getElementById("login-section");
const dashboardSection = document.getElementById("dashboard-section");
const loginForm = document.getElementById("login-form");
const usernameInput = document.getElementById("username");
const passwordInput = document.getElementById("password");
const loginError = document.getElementById("login-error");
const authNav = document.getElementById("auth-nav");
const loginBtn = document.getElementById("login-btn");
const logoutBtn = document.getElementById("logout-btn");
const currentUsernameSpan = document.getElementById("current-username");
const chatbotSendBtn = document.getElementById("send-btn");

// Lesson Details Modal Elements
const lessonDetailModal = document.getElementById("lesson-detail-modal");
const lessonDetailTitle = document.getElementById("lesson-detail-title");
const lessonDetailLanguage = document.getElementById("lesson-detail-language");
const lessonDetailContent = document.getElementById("lesson-detail-content");
const lessonDetailGrammar = document.getElementById("lesson-detail-grammar");
const lessonDetailSampleSentences = document.getElementById(
  "lesson-detail-sample-sentences"
);
const lessonDetailWordsList = document.getElementById("lesson-detail-words");
const lessonDetailWordsMeaning = document.getElementById(
  "lesson-detail-words-meaning"
);

const closeLessonModalButton = document.getElementById("lesson-close-button");
const closeExerciseModalButton = document.getElementById(
  "exercise-close-button"
);

// Exercise generation buttons moved inside lesson detail modal
const generateDialogueInitBtn = document.getElementById(
  "generate-dialogue-init-btn"
);
const generateDialogueContinueBtn = document.getElementById(
  "generate-dialogue-continue-btn"
);
const generateReadingComprehensionBtn = document.getElementById(
  "generate-reading-comprehension-btn"
);
const generateTranslationBtn = document.getElementById(
  "generate-translation-btn"
);
const lessonDetailCloseButton = document.querySelector(
  ".modal-content .lesson-detail-close-button"
);

// Dashboard Lists
const lessonsList = document.getElementById("lessons-list");
const vocabularyList = document.getElementById("vocabulary-list");
const grammarList = document.getElementById("grammar-list");
const generateLessonBtn = document.getElementById("generate-lesson-btn");

// Exercise Modal Elements
const exerciseModal = document.getElementById("exercise-modal");
const exerciseTitle = document.getElementById("exercise-title");
const exerciseContent = document.getElementById("exercise-content");
const exerciseInput = document.getElementById("exercise-input");
const exerciseAnswer = document.getElementById("exercise-answer");
const submitExerciseBtn = document.getElementById("submit-exercise-btn");
const exerciseFeedback = document.getElementById("exercise-feedback");
const feedbackText = document.getElementById("feedback-text");
const feedbackScore = document.getElementById("feedback-score");
const exerciseCloseButton = document.querySelector(".exercise-close-button");

// User profile
const prefLang = document.getElementById("preferred-lang");
const targetLang = document.getElementById("target-lang");
