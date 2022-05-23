<script>
  import Head from "../Head.svelte";
  import Navbar from "../Navbar.svelte";
  export let series = [];
</script>

<Head />
<Navbar />

<h1>All Series</h1>

<div class="container-fluid">
  <div class="album py-5 bg-light">
    <div
      class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-5 row-cols-xl-6  g-2"
    >
      {#each series as issue, i}
        <div class="col">
          <div class="card shadow-sm">
            <a href="../series/{i}">
              <img
                class="card-img-top"
                style="height: 20rem; object-fit:contain;"
                src={issue["image"]["small_url"]}
                alt={issue["image"]["small_url"]}
                preserveAspectRatio="xMidYMid slice"
              />
            </a>
            <div class="card-body">
              <h5 class="card-title" style="height: 5rem;">
                {issue["series"]}
              </h5>
              <h6 class="card-subtitle mb-2 text-muted" style="height: 2rem;">
                {issue["volume"]} - {issue["publisher"]}
              </h6>
              <div class="d-flex justify-content-between align-items-center">
                <div class="btn-group">
                  <a
                    href="../series/{i}"
                    class="btn btn-sm btn-outline-secondary">Details</a
                  >
                </div>
                <small class="text-muted"
                  >{issue["count"]}/{issue["totalcount"]}</small
                >
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
