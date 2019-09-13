<script>
    import { createEventDispatcher } from "svelte";




  import notes from "../store.js";

  export let ID;
  export let title;
  export let message;
  const dispatch = createEventDispatcher();
  const removeNote = ID => {
    fetch(`http://localhost:9000/api/notes/${ID}`, {
      method: "DELETE"
    })
      .then(res => {
        notes.removeNote(ID);
      })
      .catch(err => console.log(err));
  };
</script>

<style>
  td {
    text-align: center;
  }
</style>

<td>{ID}</td>
<td>{title}</td>
<td>{message}</td>
<td>
  <button on:click={() => dispatch('edit', ID)}>EDIT</button>
  <button on:click={() => removeNote(ID)}>X</button>
</td>
