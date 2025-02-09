import { AppBar, Toolbar, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import "./App.css";
import reactLogo from "./assets/react.svg";
import "./styles.css";
import viteLogo from "/vite.svg";

function App() {
  const [count, setCount] = useState(0);
  // console.log(appbar);
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/main")
      .then((response) => response.json())
      .then((data) => setMessage(data.message))
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <>
      <div>
        TEST
        <AppBar className={"appbar"}>
          <Toolbar>
            <Typography>{message}</Typography>
          </Toolbar>
        </AppBar>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
