import { createEffect, onMount, createSignal, For } from "solid-js";

const DataTable = () => {
  const [data, setData] = createSignal([]);
  const [page, setPage] = createSignal(1);
  const [limit, setLimit] = createSignal(20);
  const [hasMore, setHasMore] = createSignal(true);

  async function fetchData(page) {
    const url = `http://localhost:3000/houses?page=${page}&limit=${limit}`;
    const response = await fetch(url);
    const json = await response.json();

    if (json.length < limit) setHasMore(false);

    setData((oldData) => [...oldData, ...json]);
  }

  async function fetchAllData() {
    const url = "http://localhost:3000/houses";
    const response = await fetch(url);
    const json = await response.json();

    setData(json);
    setHasMore(false);
  }

  onMount(() => fetchData(page(), limit()));

  createEffect(() => {
    if (page() > 1 || limit() === -1) fetchData(page(), limit());
  });

  return (
    <div className="flex-col justify-center items-center">
      <div className="overflow-x-auto">
        <table
          className="
          table
          table-xs
          cursor-default
          "
        >
          <thead>
            <tr>
              <th></th>
              <th>Price</th>
              <th>Beds</th>
              <th>Baths</th>
              <th>Sqft</th>
              <th>LotSize</th>
              <th>LotUnit</th>
              <th>LotSqft</th>
              <th>Hty</th>
              <th>HtyPcnt</th>
              <th>Street</th>
              <th>City</th>
              <th>State</th>
              <th>Zip</th>
              <th>Link</th>
              <th>CrawlTime</th>
            </tr>
          </thead>
          <For each={data()}>
            {(item, index) => (
              <tr className="hover">
                <td>{index() + 1}</td>
                <td>{item.Price}</td>
                <td>{item.Beds}</td>
                <td>{item.Baths}</td>
                <td>{item.Sqft}</td>
                <td>{item.LotSize}</td>
                <td>{item.LotUnit}</td>
                <td>{item.LotSqft}</td>
                <td>{item.Hty}</td>
                <td>{item.HtyPcnt}</td>
                <td>{item.Street}</td>
                <td>{item.City}</td>
                <td>{item.State}</td>
                <td>{item.Zip}</td>
                <td>
                  <a href={item.Link} target="_blank">
                    Link
                  </a>
                </td>
                <td>{item.CrawlTime}</td>
              </tr>
            )}
          </For>
        </table>
      </div>
      <button
        className="btn btn-sm"
        onClick={() => setPage(page() + 1)}
        disabled={!hasMore()}
      >
        Load More
      </button>
      <button
        className="btn btn-sm"
        onClick={fetchAllData}
        disabled={!hasMore()}
      >
        Show All
      </button>
    </div>
  );
};

export default DataTable;
