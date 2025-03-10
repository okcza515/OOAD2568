
import java.util.ArrayList;
import javax.swing.*;

public class Filter extends JTextField{
	
	private ListModel listModel;
    private EditorMediator _mediator;

    public Filter(EditorMediator mediator){
        this._mediator = mediator;
    }
	
	public void setList(ListModel listModel) {
        this.listModel = listModel;
    }
	
	private void searchElements(String s) {
        if (listModel == null) {
            return;
        }

        if (s.equals("")) {
            _mediator.setElementsList(listModel);
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
        _mediator.setElementsList(listModel);
    }
}
