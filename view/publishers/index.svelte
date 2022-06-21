<script>
  import Head from "../Head.svelte";
  import Navbar from "../Navbar.svelte";
  export let publishers = [];
</script>

<Head />
<Navbar title={"Publishers"}/>

<div class="container-fluid" style="overflow-y: scroll; height: 55rem;">
  <div class="album py-2 bg-light">
    <div class="row row-cols-2 row-cols-sm-3 row-cols-md-4 row-cols-lg-6 row-cols-xl-8  g-2">
      {#each publishers as publisher, i}
        <div class="col">
          <div class="card shadow-sm">
            <a href="../publishers/{publisher["id"]}">
              <img
                class="card-img-top"
                style="height: 15rem; object-fit:contain;"
                src={publisher["image"]["small_url"]}
                alt={publisher["image"]["small_url"]}
                preserveAspectRatio="xMidYMid slice"
              />
            </a>
            <div class="card-body">
              <h5 class="card-title" style="height: 3rem;">
                {publisher["name"]}
              </h5>
              <div class="d-flex justify-content-between align-items-center">
                <div class="btn-group">
                </div>
                <small class="text-muted"
                  >Series count: <b>{publisher["series_count"]}</b></small
                >
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>

