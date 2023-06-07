const ViewSwitcher = (props) => {
  return (
    <div className="flex justify-center w-full mt-5 mb-5">
      <div className="join">
        <input
          className="join-item btn"
          type="radio"
          name="options"
          aria-label="Table"
          checked={props.viewState() === "table"}
          onChange={() => props.setViewState("table")}
        />
        <input
          className="join-item btn"
          type="radio"
          name="options"
          aria-label="Stats"
          checked={props.viewState() === "stats"}
          onChange={() => props.setViewState("stats")}
        />
      </div>
    </div>
  );
};

export default ViewSwitcher;
