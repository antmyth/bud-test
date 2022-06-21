<script>
  import Head from "../Head.svelte";
  import Navbar from "../Navbar.svelte";
  export let publisher = {};
</script>

<Head />
<Navbar title={"Series Detail"}/>

<div class="container border" style="overflow-y: scroll; height: 60rem;">
  <div class="row border">
    <div class="col border">
      <img
        style="height: 35rem; width:25rem; object-fit:contain;"
        src={publisher["image"]["original_url"]}
        alt={publisher["image"]["small_url"]}
        preserveAspectRatio="xMidYMid slice"
      />
    </div>
    <div class="col-7">
      <div class="row">
        <div class="col border bg-light">
          <h1>{publisher["name"]}</h1>
        </div>
        <div class="w-100" />
        <!-- <div class="col border bg-light">
          <h4>{publisher["volume"]}</h4>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h3>{publisher["publisher"]}</h3>
        </div>
        <div class="w-100" /> -->
        <div class="col border bg-light" style="overflow-y: scroll; height: 30rem;">
          <p>{@html publisher["description"]}</p>
        </div>
        <!-- <div class="w-100" />
        <div class="col border bg-light">
          <h5>Location:</h5>
          <i>{publisher["location"]}</i>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h6><a href={publisher["web"]}>Comic Vine Page</a></h6>
        </div> -->
        <div class="w-100" />
        <div class="col border bg-light">
          <h5>
            Series available: <b>{publisher["series_count"]}</b>
          </h5>
        </div>
      </div>
    </div>
  </div>
  <div class="row">
    <div class="col">--</div>
  </div>
  <div class="row align-items-center row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 row-cols-xl-5 g-2">
    {#each publisher["series"] as series}
    <div class="col">
      <div class="card shadow-sm">
        <a href="../series/{series["id"]}">
          <img
          class="bd-placeholder-img card-img-top"
          style="height: 15rem; object-fit:contain;"
          src={series["image"]["small_url"]}
          alt={series["image"]["small_url"]}
          preserveAspectRatio="xMidYMid slice"
          />
        </a>
        <div class="card-body">
          <h5 class="card-title" style="height: 3rem;">
            {series["series"]}
          </h5>
          <div class="d-flex justify-content-between align-items-center">
            <small class="text-muted"> {series["count"]}/{series["totalcount"]}</small>
            <small class="text-muted">{series["volume"]}</small>
          </div>
        </div>
      </div>
    </div>
    {/each}
  </div>
</div>