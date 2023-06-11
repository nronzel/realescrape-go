export async function getCount() {
  const res = await fetch("http://localhost:3000/houses/count");
  const count = await res.json();

  return count.count;
}

export async function getData(currentPage, limit = 20) {
  const url = `http://localhost:3000/houses?page=${currentPage}&limit=${limit}`;
  const res = await fetch(url);
  const data = await res.json();

  return data;
}

export async function getAllData() {
  const url = "http://localhost:3000/houses";
  const res = await fetch(url);
  const data = await res.json();

  return data;
}

export async function deleteData() {
  const url = "http://localhost:3000/cleardb";
  const res = await fetch(url, { method: "POST" });
  const json = await res.json();

  console.log(json);
}
