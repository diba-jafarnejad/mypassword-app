import React, { useState } from "react";
import Auth from "./components/Auth";
import PasswordManager from "./components/PasswordManager";

function App() {
  const [isAuthenticated, setAuthenticated] = useState(false);

  return (
    <div className="App">
      {isAuthenticated ? (
        <PasswordManager />
      ) : (
        <Auth setAuthenticated={setAuthenticated} />
      )}
    </div>
  );
}

export default App;