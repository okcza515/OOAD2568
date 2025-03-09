
import java.util.ArrayList;
import javax.swing.*;

public class Filter extends JTextField{
	
	private ListModel listModel;
    private EditorMediator mediator;

    public Filter(EditorMediator mediator) { 
        super(20);
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
        DefaultListModel<Note> listModel = new DefaultListModel<>();
        for (Note note : notes) {
            if (note.getName().contains(s)) {
                listModel.addElement(note);
            }
        }
        mediator.setElementsList(listModel);
    }
}
