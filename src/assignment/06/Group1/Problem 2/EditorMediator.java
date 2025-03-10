import javax.swing.ListModel;

public interface EditorMediator {
    void addNewNote(Note note);
    void deleteNote();
    void saveChanges();
    void setElementsList(ListModel listModel);
    void sendToFilter(ListModel listModel);
    void markNote();
    void hideElements(boolean flag);
}
