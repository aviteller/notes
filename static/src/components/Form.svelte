<script>
  import notes from "../store.js";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let ID = null;

  let title = "";
  let message = "";

  if (ID) {
    const unsubscribe = notes.subscribe(items => {
      const selectedNote = items.find(i => i.ID === ID);
      title = selectedNote.title;
      message = selectedNote.message;
    });

    unsubscribe();
  }

  const onSubmit = () => {
    let newNote = {
      title,
      message
    };

    if (!newNote.title || !newNote.message) {
      alert("please fill in all fields");
    } else {
      if (ID) {
        newNote.ID = ID
        fetch(`http://localhost:9000/api/notes/${ID}`, {
          method: "PUT",
          body: JSON.stringify(newNote),
          headers: { "Content-Type": "application/json" }
        })
          .then(res => {
            if (!res.ok) {
              throw new Error("Failed");
            }
            return res.json();
          })
          .then(data => {
            if (!data.status) {
              throw new Error(data.message);
            }
            notes.updateNote(ID, newNote);
            dispatch('cancel')
          })
          .catch(err => console.log(err));
      } else {
        fetch("http://localhost:9000/api/notes", {
          method: "POST",
          body: JSON.stringify(newNote),
          headers: { "Content-Type": "application/json" }
        })
          .then(res => {
            if (!res.ok) {
              throw new Error("Failed");
            }
            return res.json();
          })
          .then(data => {
            if (!data.status) {
              throw new Error(data.message);
            }
            notes.addNote({ ...newNote, ID: data.note.ID });
            title = "";
            message = "";
          })
          .catch(err => console.log(err));
      }
    }
  };
</script>

<style>
  .form-wrapper {
    display: flex;
    flex-direction: row;
    justify-content: left;
  }
</style>

<div class="form-wrapper">
  <div>
    <label for="title">Title</label>
    <input type="text" id="title" bind:value={title} />
  </div>
  <div>
    <label for="message">Message</label>
    <textarea id="message" cols="30" rows="3" bind:value={message} />
  </div>
  <button on:click={onSubmit}>Submit</button>
  {#if ID}
    <button on:click={() => dispatch('cancel')}>Cancel</button>
  {/if}
</div>
