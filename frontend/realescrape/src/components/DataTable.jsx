import { createEffect, onMount, createSignal, For } from "solid-js";
import useSort from "./useSort";
import SortableHeader from "./SortableHeader";

const DataTable = (props) => {
  const [page, setPage] = createSignal(1);
  const [hasMore, setHasMore] = createSignal(true);
  const [total, setTotal] = createSignal(0);
  const [sortingNeeded, setSortingNeeded] = createSignal(false);
  const sort = useSort("CrawlTime");

  async function fetchCount() {
    const response = await fetch("http://localhost:3000/houses/count");
    const json = await response.json();

    setTotal(json.count);
  }

  async function fetchData(currentPage, append = false) {
    const url = `http://localhost:3000/houses?page=${currentPage}&limit=20`;
    const response = await fetch(url);
    const json = await response.json();

    if (json.length === 0 || (currentPage - 1) * 20 + json.length >= total()) {
      setHasMore(false);
    } else {
      setHasMore(true);
    }

    if (append) {
      props.setData((oldData) => [...oldData, ...json]);
    } else {
      props.setData(json);
    }
    setSortingNeeded(true);
  }

  async function fetchAllData() {
    const url = "http://localhost:3000/houses";
    const response = await fetch(url);
    const json = await response.json();

    props.setData(json);
    setHasMore(false);
  }

  async function deleteData() {
    const url = "http://localhost:3000/cleardb";
    const response = await fetch(url, { method: "POST" });
    const json = await response.json();

    console.log(json);
  }

  onMount(() => {
    fetchCount().then(() => {
      fetchData(page());
      if (total() > 20) {
        setHasMore(true);
      }
    });
  });

  createEffect(() => {
    if (!sortingNeeded()) return;
    if (sortingNeeded()) {
      const sortedData = sort.sortByColumn(props.data());
      props.setData(sortedData);
      setSortingNeeded(false);
    }
  });

  createEffect(() => {
    if (!props.searchPerformed()) return;
    if (props.searchPerformed()) {
      total() > 20 ? setHasMore(true) : "";
      setPage(1);
      fetchCount();
      fetchData(1);
      props.onSearch(false);
    }
  });

  createEffect(() => {
    const currentPage = page();
    if (currentPage > 1 && props.data().length < total())
      fetchData(currentPage, true);
  });

  return (
    <div className="flex flex-col justify-center items-center">
      {
        <div
          className="
          stats
          stats-vertical
          lg:stats-horizontal
          shadow
          border
          border-gray-700
          mb-5
          "
        >
          <div className="stat place-items-center">
            <div className="stat-title">Total Listings</div>
            <div className="stat-value">{total()}</div>
            <div className="stat-desc"></div>
          </div>

          <div className="stat place-items-center">
            <div className="stat-title">Clear Database</div>
            <div className="stat-value">
              <button
                className="btn btn-xs btn-outline btn-error hover:bg-red-800 hover:text-white"
                onClick={deleteData}
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      }
      {props.data().length > 0 ? (
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
                  <SortableHeader
                    column="Price"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Price
                  </SortableHeader>
                  <SortableHeader
                    column="Beds"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Beds
                  </SortableHeader>
                  <SortableHeader
                    column="Baths"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Baths
                  </SortableHeader>
                  <SortableHeader
                    column="Sqft"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Sqft
                  </SortableHeader>
                  <SortableHeader
                    column="LotSize"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    LotSize
                  </SortableHeader>
                  <SortableHeader
                    column="LotUnit"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    LotUnit
                  </SortableHeader>
                  <SortableHeader
                    column="LotSqft"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    LotSqft
                  </SortableHeader>
                  <SortableHeader
                    column="Hty"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Hty
                  </SortableHeader>
                  <SortableHeader
                    column="HtyPcnt"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    HtyPcnt
                  </SortableHeader>
                  <SortableHeader
                    column="Street"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Street
                  </SortableHeader>
                  <SortableHeader
                    column="City"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    City
                  </SortableHeader>
                  <SortableHeader
                    column="State"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    State
                  </SortableHeader>
                  <SortableHeader
                    column="Zip"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    Zip
                  </SortableHeader>
                  <SortableHeader column="Link" handleSort={sort.handleSort}>
                    Link
                  </SortableHeader>
                  <SortableHeader
                    column="CrawlTime"
                    handleSort={(column) => {
                      sort.handleSort(column);
                      setSortingNeeded(true);
                    }}
                  >
                    CrawlTime
                  </SortableHeader>
                </tr>
              </thead>
              <For each={props.data()}>
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
                    <td>
                      {item.CrawlTime.split("T")[0] +
                        " " +
                        item.CrawlTime.split("T")[1].replace(/\..*$/, "")}
                    </td>
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
              className="btn btn-outline btn-sm btn-neutral"
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
