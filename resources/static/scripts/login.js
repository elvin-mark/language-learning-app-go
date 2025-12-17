// --- Authentication ---

async function loginUser(username, password) {
  loginError.textContent = ""; // Clear previous errors
  const requestBody = JSON.stringify({
    username: username,
    password: password,
  });

  try {
    const response = await fetch(`${API_BASE_URL}/auth/token`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: requestBody,
    });

    if (!response.ok) {
      const errorData = await response.json();
      // More robust error handling for different status codes if available
      const errorMessage =
        errorData.error ||
        errorData.message ||
        `Login failed. Status: ${response.status}`;
      throw new Error(errorMessage);
    }

    const authData = await response.json();
    // Assuming the API returns an array with a single token object
    if (authData) {
      localStorage.setItem(AUTH_TOKEN_KEY, authData.AccessToken);
      localStorage.setItem(USER_ID_KEY, authData.UserId);
      localStorage.setItem(USERNAME_KEY, username);
      updateUIForLoggedInState(username);
      await loadDashboardData(authData.AccessToken);
    } else {
      throw new Error(
        "Invalid response from authentication server. No token found."
      );
    }
  } catch (error) {
    console.error("Login error:", error);
    loginError.textContent = error.message;
  }
}

function logoutUser() {
  localStorage.removeItem(AUTH_TOKEN_KEY);
  localStorage.removeItem(USER_ID_KEY);
  localStorage.removeItem(USERNAME_KEY);
  updateUIForLoggedOutState();
  clearDashboardData();
}
