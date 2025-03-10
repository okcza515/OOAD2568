// 65070501049 Roodfan Maimahad

import javax.swing.*;

public interface Mediator {
    void addNewNote(Note note);
    void deleteNote();
    void getInfoFromList(Note note);
    void saveChanges();
    void markNote();
    void clear();
    void sendToFilter(ListModel listModel);
    void setElementsList(ListModel list);
    void createElement(Element element);
    void hideElements(boolean flag);
    void createGUI();
}