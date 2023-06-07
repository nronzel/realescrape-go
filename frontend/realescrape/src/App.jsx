import Header from "./components/Header.jsx";
import SearchBar from "./components/SearchBar.jsx";
import DataTable from "./components/DataTable.jsx";
import { createSignal } from "solid-js";
import ViewSwitcher from "./components/ViewSwitcher.jsx";

function App() {
  const [searchPerformed, setSearchPerformed] = createSignal(false);
  const [searchQuery, setSearchQuery] = createSignal("");
  const [viewState, setViewState] = createSignal("table");

  return (
    <div className="main">
      <Header />
      <SearchBar onSearch={setSearchPerformed} onQuery={setSearchQuery} />
      <ViewSwitcher viewState={viewState} setViewState={setViewState} />
      {viewState() === "table" ? (
        <DataTable
          searchPerformed={searchPerformed}
          searchQuery={searchQuery}
        />
      ) : (
        <div>stat view</div>
      )}
    </div>
  );
}

export default App;
