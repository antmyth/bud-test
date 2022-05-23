<script>
  export let stories = []
</script>

<h1>Story Index</h1>

<table border="1" cellpadding="10">
  {#if stories.length > 0}
    <thead>
      {#each Object.keys(stories[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each stories as story}
    <tr>
      {#each Object.keys(story) as key}
        {#if key === "id"}
          <td><a href={`/${story[key]}`}>{story[key]}</a></td>
        {:else}
          <td>{story[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<style>
  table {
    border-collapse: collapse;
  }
</style>
