<script>
  export let issue = {}
</script>

<h1>Showing Issue</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(issue) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(issue) as key}
      <td>{issue[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href="../issues">Back</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
