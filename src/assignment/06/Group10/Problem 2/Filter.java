
import java.util.ArrayList;
import javax.swing.*;

public class Filter extends JTextField{
	
	private ListModel listModel;
	
	public void setList(ListModel listModel) {
        this.listModel = listModel;
    }
	
	private void searchElements(String s) {
        if (listModel == null) {
            return;
        }

        if (s.equals("")) {
            Editor.setElementsList(listModel);
            return;
        }

        ArrayList<Note> notes = new ArrayList<>();
        for (int i = 0; i < listModel.getSize(); i++) {
            notes.add((Note) listModel.getElementAt(i));
        }
        DefaultListModel<Note> listModel = new DefaultListModel<>();
        for (Note note : notes) {
            if (note.getName().contains(s)) {
                listModel.addElement(note);
            }
        }
        Editor.setElementsList(listModel);
    }
}
