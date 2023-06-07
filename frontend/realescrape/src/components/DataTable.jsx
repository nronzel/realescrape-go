import { createEffect, onMount, createSignal, For } from "solid-js";

const DataTable = (props) => {
  const [data, setData] = createSignal([]);
  const [page, setPage] = createSignal(1);
  const [hasMore, setHasMore] = createSignal(true);

  async function fetchData(currentPage, append = false) {
    const url = `http://localhost:3000/houses?page=${currentPage}&limit=20`;
    const response = await fetch(url);
    const json = await response.json();

    if (json.length === 0) setHasMore(false);

    if (append) {
      setData((oldData) => [...oldData, ...json]);
    } else {
      setData(json);
    }
  }

  async function fetchAllData() {
    const url = "http://localhost:3000/houses";
    const response = await fetch(url);
    const json = await response.json();

    setData(json);
    // setHasMore(false);
  }

  onMount(() => fetchData(page()));

  createEffect(() => {
    if (props.searchPerformed()) {
      setPage(1);
      fetchData(1);
      props.onSearch(false);
    }
  });

  createEffect(() => {
    const currentPage = page();
    if (currentPage > 1) fetchData(currentPage, true);
  });

  return (
    <div className="flex flex-col justify-center items-center">
      {data().length > 0 ? (
        <>
          <div className="overflow-x-auto w-11/12">
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
              <tfoot>
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
              </tfoot>
            </table>
          </div>
          <div className="flex gap-4 mt-5 mb-5 justify-center w-full">
            <button
              className="btn btn-sm btn-secondary"
              onClick={() => setPage(page() + 1)}
              disabled={!hasMore()}
            >
              Load More
            </button>
            <button
              className="btn btn-sm btn-accent"
              onClick={fetchAllData}
              disabled={!hasMore()}
            >
              Show All
            </button>
          </div>
        </>
      ) : (
        <div className="flex justify-center mt-10">
          <p className="text-lg">
            No results in database. Perform a search above.
          </p>
        </div>
      )}
    </div>
  );
};

export default DataTable;
