<script>
  import Head from "../Head.svelte";
  import CardIssue from "../issues/CardIssue.svelte";
  import Navbar from "../Navbar.svelte";
  export let series = {};
</script>

<Head />
<Navbar />

<div class="container border">
  <div class="row border">
    <div class="col border">
      <img
        style="height: 35rem; width:25rem; object-fit:contain;"
        src={series["image"]["original_url"]}
        alt={series["image"]["small_url"]}
        preserveAspectRatio="xMidYMid slice"
      />
    </div>
    <div class="col-6">
      <div class="row">
        <div class="col border bg-light">
          <h1>{series["series"]}</h1>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h4>{series["volume"]}</h4>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h3>{series["publisher"]}</h3>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <p>{@html series["description"]}</p>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h5>Location:</h5>
          <i>{series["location"]}</i>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h6><a href={series["web"]}>Comic Vine Page</a></h6>
        </div>
        <div class="w-100" />
        <div class="col border bg-light">
          <h5>
            Issues available: <b>{series["count"]}</b> of
            <b>{series["totalcount"]}</b>
          </h5>
        </div>
      </div>
    </div>
  </div>
  <div class="row">
    <div class="col"><h4>Availalbe Issues</h4></div>
  </div>
  <div class="row align-items-center row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 row-cols-xl-5 g-2">
    {#each series["issues"] as issue, i}
      <CardIssue {issue} />
    {/each}
  </div>
</div>

<br />

<a href="../series">Back</a>
