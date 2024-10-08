import React from "react";
import PasswordManager from "./components/PasswordManager";
import Auth from "./components/Auth";

function App() {
  const [isAuthenticated, setAuthenticated] = React.useState(false);

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

