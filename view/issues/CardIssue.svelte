<script>
    export let issue;
  </script>
      <div class="col">
        <div class="card shadow-sm">
          <img
            class="bd-placeholder-img card-img-top"
            style="height: 15rem; object-fit:contain;"
            src={issue["image"]["small_url"]}
            alt={issue["image"]["small_url"]}
            preserveAspectRatio="xMidYMid slice"
          />
          <div class="card-body">
            <p class="card-text"  style="height: 2rem;">
              {issue["number"]} - {issue["title"]}
            </p>
            <div class="d-flex justify-content-between align-items-center">
              <small class="text-muted">{issue["volume"]}</small>
            </div>
          </div>
        </div>
      </div>