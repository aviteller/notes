<script>
  import NoteItem from './NoteItem.svelte'
  export let notes;

</script>


<style>
.notes-table {width:100%}
</style>

<table class="notes-table">
  <thead>
    <tr>
      <th>ID</th>
      <th>Title</th>
      <th>Message</th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    {#each notes as note (note.ID)}
      <tr>
        <NoteItem ID={note.ID} title={note.title} message={note.message} on:edit/>
      </tr>
    {/each}
  </tbody>
</table>