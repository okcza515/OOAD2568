import java.util.ArrayList;

class List {
    private java.util.List<Note> notes = new ArrayList<>();
    private Note selectedNote;

    public void addElement(Note note) {
        notes.add(note);
    }

    public Note getSelectedElement() {
        return selectedNote;
    }

    public void removeSelectedElement() {
        if (selectedNote != null) {
            notes.remove(selectedNote);
            selectedNote = null;
        }
    }

    public void highlightSelectedElement() {
        if (selectedNote != null) {
            System.out.println("Highlighting note: " + selectedNote.getName());
        }
    }
}
