import javax.swing.*;
import java.util.ArrayList;

public class Filter extends JTextField {

    private Editor mediator;
    private ListModel listModel;

    public Filter(Editor mediator) {
        super();
        this.mediator = mediator;
    }

    public void setList(ListModel listModel) {
        this.listModel = listModel;
    }

    private void searchElements(String s) {
        if (listModel == null) {
            return;
        }

        if (s.equals("")) {
            mediator.setElementsList(listModel);
            return;
        }

        ArrayList<Note> notes = new ArrayList<>();
        for (int i = 0; i < listModel.getSize(); i++) {
            notes.add((Note) listModel.getElementAt(i));
        }
        DefaultListModel<Note> newModel = new DefaultListModel<>();
        for (Note note : notes) {
            if (note.getName().contains(s)) {
                newModel.addElement(note);
            }
        }
        mediator.setElementsList(newModel);
    }
}