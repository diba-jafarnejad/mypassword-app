import React, { useState, useEffect } from "react";
import axios from "../services/api";

interface PasswordEntry {
  id: number;
  account: string;
  username: string;
  password: string;
}

const PasswordManager: React.FC = () => {
  const [passwords, setPasswords] = useState<PasswordEntry[]>([]);

  useEffect(() => {
    const fetchPasswords = async () => {
      try {
        const response = await axios.get<PasswordEntry[]>('/passwords');
        setPasswords(response.data);
      }
      catch (error) {
        console.error('Error fetching passwords', error);
      }
      
    };
    fetchPasswords();
  }, []);

  return (
    <div>
      <h2>Stored Passwords</h2>
      <ul>
        {passwords.map((entry, index) => (
          <li key={index}>
            {entry.account} - {entry.username} - {entry.password}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PasswordManager;