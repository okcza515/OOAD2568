import javax.swing.ListModel;

public interface EditorMediator {
    void addNote();
    void deleteNote();
    void saveNote();
    void selectNote(Note note);
    void clearNote();
    void markNote();
    void setElementsList(ListModel listModel);
    void sendToFilter(ListModel listModel);
    void hideElements(boolean flag);
}
