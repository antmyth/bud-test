<script>
  import Head from "../Head.svelte";
  import Navbar from "../Navbar.svelte";
  export let series = [];
</script>

<Head />
<Navbar title={"Series"}/>
<div class="container-fluid">
  <nav class="navbar navbar-expand-lg navbar-light bg-secondary">
    <ul class="pagination">
      <li class="page-item"><a class="page-link" href="../series?page=0&size=300">1</a></li>
      <li class="page-item"><a class="page-link" href="../series?page=1&size=300">2</a></li>
      <li class="page-item"><a class="page-link" href="../series?page=2&size=300">3</a></li>
    </ul>
  </nav>
</div>



<div class="container-fluid" style="overflow-y: scroll; height: 50rem;">
  <div class="album py-2 bg-light">
    <div
      class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-5 row-cols-xl-6  g-2"
    >
      {#each series as ser, i}
        <div class="col">
          <div class="card shadow-sm">
            <a href="../series/{ser["id"]}">
              <img
                class="card-img-top"
                style="height: 20rem; object-fit:contain;"
                src={ser["image"]["small_url"]}
                alt={ser["image"]["small_url"]}
                preserveAspectRatio="xMidYMid slice"
              />
            </a>
            <div class="card-body">
              <h5 class="card-title" style="height: 5rem;">
                {ser["series"]}
              </h5>
              <h6 class="card-subtitle mb-2 text-muted" style="height: 2rem;">
                {ser["volume"]} - {ser["publisher"]}
              </h6>
              <div class="d-flex justify-content-between align-items-center">
                <div class="btn-group">
                  <a
                    href="../series/{ser["id"]}"
                    class="btn btn-sm btn-outline-secondary">Details</a
                  >
                </div>
                <small class="text-muted"
                  >{ser["count"]}/{ser["totalcount"]}</small
                >
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
