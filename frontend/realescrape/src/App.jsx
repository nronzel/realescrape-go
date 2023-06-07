import Header from "./components/Header.jsx";
import SearchBar from "./components/SearchBar.jsx";
import DataTable from "./components/DataTable.jsx";
import { createSignal } from "solid-js";

function App() {
  const [searchPerformed, setSearchPerformed] = createSignal(false);
  const [searchQuery, setSearchQuery] = createSignal("");

  return (
    <div className="main">
      <Header />
      <SearchBar onSearch={setSearchPerformed} onQuery={setSearchQuery} />
      <DataTable searchPerformed={searchPerformed} searchQuery={searchQuery} />
    </div>
  );
}

export default App;
