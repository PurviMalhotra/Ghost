import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/ping")
      .then((res) => res.text())
      .then((data) => setMessage(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div>
      <h1>Ghost Test</h1>
      <img
        src="/ghost-assets/bread.jpg"
        width="400"
      />
    </div>
  );
}

export default App;