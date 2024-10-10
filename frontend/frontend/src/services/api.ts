import axios from "axios";

// Create an axios instance with a base URL for your API
export const api = axios.create({
  baseURL: "http://localhost:8080", // Your backend URL
  headers: {
    "Content-Type": "application/json", // Default content type for JSON requests
  },
});

// The login function that makes a POST request with correct headers
export const login = (username: string, password: string) => {
  return api.post(
    '/auth', 
    { username, password },  // Payload for the login request
    {
      headers: {
        'Content-Type': 'application/json', // Explicitly set content type for the POST request
      },
    }
  );
};
