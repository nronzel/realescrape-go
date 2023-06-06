const SearchBar = () => {
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
        />
        <label className="label">
          <span className="label-text-alt">
            ex. Miami FL, 90210, San-Francisco CA
          </span>
        </label>
      </div>
      <input type="button" value="Search" className="btn btn-primary btn-sm" />
    </div>
  );
};

export default SearchBar;
