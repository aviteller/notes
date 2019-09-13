<script>
  import Header from "./components/Header.svelte";
  import NoteTable from "./components/NoteTable.svelte";
  import Form from "./components/Form.svelte";
  import notes from "./store.js";
  let editedId = null;
  let editMode = null;
  let isLoading = false;
  const getNotes = () => {
    fetch(`http://localhost:9000/api/notes`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching meetups");
        }
        return res.json();
      })
      .then(data => {
        if (!data.status) {
          error = data;
        } else {
          const loadedNotes = [];
          for (const key in data.data) {
            loadedNotes.push({
              ...data.data[key]
            });
          }
          notes.setNotes(loadedNotes);
          isLoading = false;
        }
      })
      .catch(err => {
        error = err;
        isLoading = false;
      });
  };
  getNotes();

  const startEdit = e => {
    editMode = "edit";
    editedId = e.detail;
	};
	
	const cancelEdit = () =>{
		editMode = null
	}
</script>

<style>
  main {
    margin-top: 5rem;
  }
</style>

<Header />
<main>
  {#if editMode == 'edit'}
    <Form ID={editedId} on:cancel={cancelEdit}/>
  {:else}
    <Form />
  {/if}
  <NoteTable notes={$notes} on:edit={startEdit} />
</main>
