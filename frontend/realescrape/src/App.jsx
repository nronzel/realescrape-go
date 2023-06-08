import Header from "./components/Header.jsx";
import SearchBar from "./components/SearchBar.jsx";
import DataTable from "./components/DataTable.jsx";
import { createSignal } from "solid-js";
import ViewSwitcher from "./components/ViewSwitcher.jsx";

function App() {
  const [searchPerformed, setSearchPerformed] = createSignal(false);
  const [viewState, setViewState] = createSignal("table");
  const [data, setData] = createSignal([]);

  return (
    <div className="main">
      <Header />
      <SearchBar setData={setData} onSearch={setSearchPerformed} />
      <ViewSwitcher
        viewState={viewState}
        setViewState={setViewState}
        setSearchPerformed={setSearchPerformed}
      />
      {viewState() === "table" ? (
        <DataTable
          data={data}
          setData={setData}
          onSearch={setSearchPerformed}
          searchPerformed={searchPerformed}
        />
      ) : (
        <p>Stats page</p>
      )}
    </div>
  );
}

export default App;
