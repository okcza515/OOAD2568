import javax.swing.*;

public interface EditorMediator {
    void getInfoFromList(Note note);
    void clear();
    void hideElements(boolean flag);
    void addNewNote(Note note);
    void sendToFilter(ListModel listModel);
    void setElementsList(ListModel listM);
    void markNote();
    void deleteNote();
    void saveChanges();
    void createGUI();
}
