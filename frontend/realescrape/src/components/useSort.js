import { createSignal } from "solid-js";

export default function useSort(column = "") {
  const [sortConfig, setSortConfig] = createSignal({
    column,
    direction: "asc",
  });

  const sortByColumn = (array) => {
    const sortDetails = sortConfig();
    return array.slice().sort((a, b) => {
      if (a[sortDetails.column] < b[sortDetails.column]) {
        return sortDetails.direction === "asc" ? -1 : 1;
      }
      if (a[sortDetails.column] > b[sortDetails.column]) {
        return sortDetails.direction === "asc" ? 1 : -1;
      }
      return 0;
    });
  };

  const handleSort = (column) => {
    setSortConfig((prevSortConfig) => {
      const newSortConfig =
        prevSortConfig.column === column
          ? {
              ...prevSortConfig,
              direction: prevSortConfig.direction === "asc" ? "desc" : "asc",
            }
          : { column, direction: "asc" };
      return newSortConfig;
    });
  };

  return { sortConfig, setSortConfig, sortByColumn, handleSort };
}
