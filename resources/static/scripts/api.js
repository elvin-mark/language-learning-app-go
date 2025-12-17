// --- API Service Functions ---

async function apiFetch(endpoint, options = {}) {
  const token = localStorage.getItem(AUTH_TOKEN_KEY);
  if (!token) {
    // If not authenticated, redirect to login or show message
    logoutUser(); // This will update UI and clear data
    throw new Error("Not authenticated. Please log in.");
  }

  const defaultHeaders = {
    "Content-Type": "application/json",
    Authorization: `Basic ${token}`,
  };

  const config = {
    ...options,
    headers: {
      ...defaultHeaders,
      ...(options.headers || {}),
    },
  };

  // Ensure body is stringified if it's an object and Content-Type is JSON
  if (
    config.body &&
    typeof config.body === "object" &&
    config.headers["Content-Type"] === "application/json"
  ) {
    config.body = JSON.stringify(config.body);
  }

  try {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, config);

    if (!response.ok) {
      let errorMsg = "An unexpected API error occurred";
      try {
        const errorData = await response.json();
        // Try to extract a meaningful error message
        errorMsg =
          errorData.error || errorData.message || JSON.stringify(errorData);
      } catch (e) {
        // If response is not JSON, use status text
        errorMsg = response.statusText;
      }

      // Handle 401 Unauthorized specifically
      if (response.status === 401) {
        console.error("API Error 401: Not Authenticated");
        logoutUser();
        alert("Your session has expired. Please log in again.");
        throw new Error("Authentication failed. Please log in again.");
      }

      throw new Error(`API Error: ${response.status} - ${errorMsg}`);
    }

    // Handle cases where response might be empty (e.g., 204 No Content)
    if (response.status === 204) {
      return null;
    }

    return response.json();
  } catch (error) {
    console.error(`API Fetch Error for ${endpoint}:`, error);
    throw error; // Re-throw to be caught by calling functions
  }
}

async function getUserProfile() {
  const user = await apiFetch("/user/profile");
  if (user) {
    return user; // Return the first user object
  } else {
    console.warn("User profile not found or API returned empty array.");
    return null;
  }
}

async function getLessons() {
  const lessons = await apiFetch("/resources/lessons");
  return lessons || []; // Ensure we always return an array
}

async function getVocabulary() {
  const words = await apiFetch("/resources/words");
  return words || []; // Ensure we always return an array
}

async function getGrammar() {
  const grammar = await apiFetch("/resources/grammar");
  return grammar || []; // Ensure we always return an array
}

async function generateNewLesson() {
  // The API POST /resources/lessons/generate doesn't take a body.
  // it generates based on user's current status implicitly.
  return apiFetch("/resources/lessons/generate", { method: "POST" });
}
