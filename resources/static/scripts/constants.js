const API_BASE_URL = "http://localhost:8081";
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
const startExerciseBtn = document.getElementById("start-exercise-btn");
const closeModalButton = document.querySelector(".modal-content .close-button");

// Dashboard Lists
const lessonsList = document.getElementById("lessons-list");
const vocabularyList = document.getElementById("vocabulary-list");
const grammarList = document.getElementById("grammar-list");
const generateLessonBtn = document.getElementById("generate-lesson-btn");

// User profile
const prefLang = document.getElementById("preferred-lang");
const targetLang = document.getElementById("target-lang");
