import React, { useState } from "react";
import axios from "../services/api";

interface AuthProps {
  setAuthenticated: (isAuthenticated: boolean) => void;
}

const Auth: React.FC<AuthProps> = ({ setAuthenticated }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const login = async () => {
    try {
      const response = await axios.post("/auth", { username, password });
      if (response.status === 200) {
        setAuthenticated(true);
      }
    } catch (error) {
      alert("Invalid credentials");
    }
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Username"
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={login}>Login</button>
    </div>
  );
};

export default Auth;
