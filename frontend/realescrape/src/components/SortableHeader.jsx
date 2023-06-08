const SortableHeader = (props) => {
  const { column, handleSort, children } = props;

  return (
    <th>
      <button onClick={() => handleSort(column)}>{children}</button>
    </th>
  );
};

export default SortableHeader;
