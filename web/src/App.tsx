import { useEffect, useState } from "react";

function App() {
  const [assets, setAssets] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/assets")
      .then((res) => res.json())
      .then((data) => setAssets(data))
      .catch(console.error);
  }, []);

  return (
    <div>
      <h1>Ghost Assets</h1>

      {assets.map((asset) => (
        <div key={asset.name}>
          <p>{asset.name}</p>
          <p>{asset.size} bytes</p>
          <p>{asset.hash.slice(0, 12)}...</p>

          <img
            src={`/ghost-assets/${asset.name}`}
            width={300}
          />
      </div>
))}
    </div>
  );
}

export default App;