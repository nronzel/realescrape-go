import { createSignal } from "solid-js";

const SearchBar = (props) => {
  const [location, setLocation] = createSignal("");

  const searchLocation = async () => {
    try {
      props.setLoadingState(true);
      const encodedLocation = encodeURIComponent(location());
      const url = `http://localhost:3000/scrape/${encodedLocation}`;
      const response = await fetch(url, { method: "POST" });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error:", error);
      props.setLoadingState(false);
    }
  };

  return (
    <div className="flex justify-center items-center gap-4">
      <div className="form-control w-full max-w-xs">
        <label className="label">
          <span className="label-text">Search for location: </span>
        </label>
        <input
          type="text"
          placeholder=""
          className="input input-bordered w-full max-w-xs outline-none"
          onInput={(e) => setLocation(e.target.value)}
        />
        <label className="label">
          <span className="label-text-alt">
            ex. Miami FL, 90210, San-Francisco CA
          </span>
        </label>
      </div>
      <input
        type="button"
        value="Search"
        className="btn btn-primary btn-sm"
        onClick={searchLocation}
      />
    </div>
  );
};

export default SearchBar;
