export async function getCount() {
  const res = await fetch("http://localhost:3000/houses/count");
  const json = await res.json();

  return json.count;
}

export async function getData(currentPage) {
  const url = `http://localhost:3000/houses?page=${currentPage}&limit=20`;
  const res = await fetch(url);
  const json = await res.json();

  return json;
}

export async function deleteData() {
    const url = "http://localhost:3000/cleardb";
    const res = await fetch(url, { method: "POST" });
    const json = await res.json();

    console.log(json)
}
