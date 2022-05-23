<script>
  import Navbar from "../Navbar.svelte";
  import Head from "../Head.svelte";
  export let issues = [];
</script>

<Head />
<Navbar />

<div class="album py-5 bg-light">
  <div class="container">
    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-5 row-cols-xl-6 g-1">
      {#each issues as issue}
        <div class="col">
          <div class="card shadow-sm">
            <img
              class="bd-placeholder-img card-img-top"
              height="350"
              src={issue["image"]["small_url"]}
              alt={issue["image"]["small_url"]}
              preserveAspectRatio="xMidYMid slice"
            />
            <div class="card-body">
              <p class="card-text">
                {issue["series"]}
              </p>
              <p class="card-text">
                {issue["number"]} - {issue["title"]}
              </p>
              <div class="d-flex justify-content-between align-items-center">
                <div class="btn-group">
                  <button type="button" class="btn btn-sm btn-outline-secondary"
                    >View</button
                  >
                </div>
                <small class="text-muted">{issue["volume"]}</small>
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
