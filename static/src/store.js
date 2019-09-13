import { writable } from "svelte/store";

const notes = writable([]);

const customnotesStore = {
  subscribe: notes.subscribe,
  setNotes: noteArray => {
    notes.set(noteArray);
  },
  addNote: noteData => {
    const newnote = {
      ...noteData
    };
    notes.update(items => [...items, newnote]);
  },
  updateNote: (id, noteData) => {
    notes.update(items => {
      const noteIndex = items.findIndex(m => m.ID === id);
      const updatednote = { ...items[noteIndex], ...noteData };
      const updatednotes = [...items];
      updatednotes[noteIndex] = updatednote;
      return updatednotes;
    });
  },
  removeNote: id => {
    notes.update(items => {
      return items.filter(i => i.ID !== id);
    });
  }
};

export default customnotesStore;
