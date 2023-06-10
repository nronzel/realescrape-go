import Header from "./components/Header.jsx";
import SearchBar from "./components/SearchBar.jsx";
import DataTable from "./components/DataTable.jsx";
import { createSignal } from "solid-js";
import ViewSwitcher from "./components/ViewSwitcher.jsx";

function App() {
  const [viewState, setViewState] = createSignal("table");
  const [data, setData] = createSignal([]);
  const [loadingState, setLoadingState] = createSignal(false);

  return (
    <div className="main">
      <Header />
      <SearchBar setLoadingState={setLoadingState} />
      <ViewSwitcher viewState={viewState} setViewState={setViewState} />
      {viewState() === "table" ? (
        <DataTable
          data={data}
          setData={setData}
          setLoadingState={setLoadingState}
          loadingState={loadingState}
        />
      ) : (
        <p>Stats page</p>
      )}
    </div>
  );
}

export default App;
